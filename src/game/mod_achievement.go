package game

import "sync"

const (
	ACHIEVEMENT_NOT_DONE = 0
	ACHIEVEMENT_DONE     = 1
)

type AchievementInfo struct {
	AchievementId int
	State         int
}

type ModAchievement struct {
	Achievements map[int]*AchievementInfo
	Locker       *sync.RWMutex
}

func (self *ModAchievement) IsAchievementDone(achievementId int) bool {
	task, ok := self.Achievements[achievementId]
	if !ok {
		return false
	}
	return task.State == ACHIEVEMENT_DONE
}
