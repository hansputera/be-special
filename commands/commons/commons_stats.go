package commons

import (
	"context"
	"fmt"
	"runtime"

	"github.com/hansputera/be-special/lib"
	"github.com/hansputera/be-special/types"
	"go.mau.fi/whatsmeow/binary/proto"
	"github.com/klauspost/cpuid/v2"
)

func StatsCommand() *types.Command {
	return &types.Command{
		Name:        "stats",
		Aliases:     []string{"statistics", "st"},
		Description: "Show system and the bot statistics",
		Action: func(ctx *lib.MessageContext) error {
			mem := &runtime.MemStats{}
			runtime.ReadMemStats(mem)

			ctx.Client.SendMessage(context.Background(), ctx.Info.Chat, &proto.Message{
				ExtendedTextMessage: &proto.ExtendedTextMessage{
					ContextInfo: &proto.ContextInfo{
						StanzaId:      &ctx.Info.ID,
						QuotedMessage: ctx.Raw,
					},
					Text: lib.ToStringPointer(fmt.Sprintf(
						"CPU: %s %d vCPUs\nProgram malloc: %d MiB\nProgram malloc total: %d MiB\nGolang GC: %d\nSystem usage: %d MiB",
						cpuid.CPU.BrandName,
						runtime.NumCPU(),
						(mem.Alloc / 1024 / 1024),
						(mem.TotalAlloc / 1024 / 1024),
						mem.NumGC,
						(mem.Sys / 1024 / 1024),
					)),
				},
			})
			return nil
		},
	}
}
