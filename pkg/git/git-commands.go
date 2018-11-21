package git

import (
	"strings"
	"github.com/isacikgoz/gitbatch/pkg/command"
	"github.com/isacikgoz/gitbatch/pkg/utils"
)


// UpstreamDifferenceCount checks how many pushables/pullables there are for the
// current branch
func UpstreamDifferenceCount(repoPath string) (string, string) {
	args := []string{"rev-list", "@{u}..HEAD", "--count"}
	pushableCount, err := command.RunCommandWithOutput(repoPath, "git", args)
	if err != nil {
		return "?", "?"
	}
	args = []string{"rev-list", "HEAD..@{u}", "--count"}
	pullableCount, err := command.RunCommandWithOutput(repoPath, "git", args)
	if err != nil {
		return "?", "?"
	}
	return strings.TrimSpace(pushableCount), strings.TrimSpace(pullableCount)
}

func CurrentBranchName(repoPath string) (string, error) {
	args := []string{"symbolic-ref", "--short", "HEAD"}
	branchName, err := command.RunCommandWithOutput(repoPath, "git", args)
	if err != nil {
		args = []string{"rev-parse", "--short", "HEAD"}
		branchName, err = command.RunCommandWithOutput(repoPath, "git", args)
		if err != nil {
			return "", err
		}
	}
	return utils.TrimTrailingNewline(branchName), nil
}

func (entity *RepoEntity) IsClean() (bool, error) {
	worktree, err := entity.Repository.Worktree()
	if err != nil {
		return true, nil
	}
	status, err := worktree.Status()
	if err != nil {
		return status.IsClean(), nil
	}
	return false, nil
}