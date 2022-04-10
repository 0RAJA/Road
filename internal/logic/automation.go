package logic

import (
	"context"
	"github.com/0RAJA/Road/internal/global"
	"time"
)

func AutoEndurancePostStar(ctx context.Context) {
	ticker := time.NewTicker(global.AllSetting.Rule.EndurancePostStarTime)
	for {
		select {
		case <-ticker.C:
			EndurancePostStar(ctx)
		}
	}
}

func AutoEnduranceViews(ctx context.Context) {
	ticker := time.NewTicker(global.AllSetting.Rule.EnduranceViewTime)
	for {
		select {
		case <-ticker.C:
			EnduranceViews(ctx)
		}
	}
}

func AutoEndurancePostViews(ctx context.Context) {
	ticker := time.NewTicker(global.AllSetting.Rule.EndurancePostViews)
	for {
		select {
		case <-ticker.C:
			EndurancePostGrowViews(ctx)
		}
	}
}

func Automation(ctx context.Context) {
	go AutoEndurancePostStar(ctx)
	go AutoEnduranceViews(ctx)
	go AutoEndurancePostViews(ctx)
}
