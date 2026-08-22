package formatting

import (
	"fmt"
	"testing"

	"github.com/bwmarrin/discordgo"
)

func makeTestUsers(n int) []*discordgo.Member {
	users := make([]*discordgo.Member, n)

	for i := range users {
		users[i] = &discordgo.Member{
			User: &discordgo.User{
				Username: "test-user",
			},
		}
	}

	return users
}

func BenchmarkGetDisplayNamesAppend(b *testing.B) {
	for _, size := range []int{1, 10, 100, 1000, 10000} {
		b.Run(fmt.Sprintf("N=%d", size), func(b *testing.B) {
			users := makeTestUsers(size)

			for b.Loop() {
				_ = GetDisplayNamesAppend(users)
			}
		})
	}
}

func BenchmarkGetDisplayNamesPreallocated(b *testing.B) {
	for _, size := range []int{1, 10, 100, 1000, 10000} {
		b.Run(fmt.Sprintf("N=%d", size), func(b *testing.B) {
			users := makeTestUsers(size)

			for b.Loop() {
				_ = GetDisplayNamesPreallocated(users)
			}
		})
	}
}