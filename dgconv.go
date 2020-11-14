package dgconv

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
)

//Getuser returns user's ID from ID, mention, username, and nickname.
func Getuser(s *discordgo.Session, user string) (id string) {
	var namediscrim string
	var name string
	var nick string
	switch {
	case regexp.MustCompile(`[0-9]+`).MatchString(user):
		_, check := s.User(user)
		if check != nil {
		} else {
			id = user
			break
		}
		fallthrough
	case regexp.MustCompile(`<@![0-9]+>`).MatchString(user):
		_, check := s.User(user[3 : len(user)-1])
		if check != nil {
		} else {
			id = user[3 : len(user)-1]
			break
		}
		fallthrough
	default:
		for _, guild := range s.State.Guilds {
			for _, mem := range guild.Members {
				if user == mem.User.Username+"#"+mem.User.Discriminator {
					if len(namediscrim) == 0 {
						namediscrim = mem.User.ID
					}
				} else if user == mem.User.Username {
					if len(name) == 0 {
						name = mem.User.ID
					}
				} else if user == mem.Nick {
					if len(nick) == 0 {
						nick = mem.User.ID
					}
				}
				if len(namediscrim) != 0 {
					break
				}
			}
			if len(namediscrim) != 0 {
				break
			}
		}
		if len(namediscrim) != 0 {
			id = namediscrim
		} else if len(name) != 0 {
			id = name
		} else if len(nick) != 0 {
			id = nick
		}
	}
	return
}

//Getrole returns role's id from ID, mention, name.
func Getrole(s *discordgo.Session, m *discordgo.MessageCreate, role string) (id string) {
	switch {
	case regexp.MustCompile(`[0-9]+`).MatchString(role):
		_, check := s.State.Role(m.GuildID, role)
		if check != nil {
		} else {
			id = role
			break
		}
		fallthrough
	case regexp.MustCompile(`<@&[0-9]+>`).MatchString(role):
		_, check := s.State.Role(m.GuildID, role[3:len(role)-1])
		if check != nil {
		} else {
			id = role[3 : len(role)-1]
			break
		}
		fallthrough
	default:
		g, _ := s.State.Guild(m.GuildID)
		for _, rol := range g.Roles {
			if role == rol.Name {
				id = rol.ID
				break
			}
		}
	}
	return
}

//Getchannel returns channel's if from ID, mantion, name.
func Getchannel(s *discordgo.Session, channel string) (id string) {
	switch {
	case regexp.MustCompile(`[0-9]+`).MatchString(channel):
		_, check := s.State.Channel(channel)
		if check != nil {
		} else {
			id = channel
			break
		}
		fallthrough
	case regexp.MustCompile(`<#[0-9]+>`).MatchString(channel):
		_, check := s.State.Channel(channel[2 : len(channel)-1])
		if check != nil {
		} else {
			id = channel[2 : len(channel)-1]
			break
		}
		fallthrough
	default:
		for _, guild := range s.State.Guilds {
			for _, ch := range guild.Channels {
				if channel == ch.Name {
					id = ch.ID
					break
				}
			}
			if len(id) != 0 {
				break
			}
		}
	}
	return
}
