package view

import "github.com/isaquecsilva/whoislive/src/domain"

type ChannelsView struct {
	Title string
	domain.Channels
}

func NewChannelsView(title string, channels domain.Channels) ChannelsView {
	return ChannelsView{
		Title:    title,
		Channels: channels,
	}
}
