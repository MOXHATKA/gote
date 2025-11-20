package main

import (
	"context"
	"fmt"
	"gote/internal/utils/ctx"
	"gote/pkg/types"
)

func StartHandler(ctx context.Context, update types.Update) {
	fmt.Println("Я сказала стартуем!")
}

func MessageHandler(ctx ctx.CustomContext, u types.Update) {
}
