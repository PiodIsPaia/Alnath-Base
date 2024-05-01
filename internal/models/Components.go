package models

import (
	"fmt"
	"sync"

	"github.com/bwmarrin/discordgo"
)

type ComponentRegistry struct {
	sync.RWMutex
	components map[string]Component
}

type Component struct {
	CustomID string
	Type     discordgo.ComponentType
	Run      func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func NewComponentRegistry() *ComponentRegistry {
	return &ComponentRegistry{
		components: make(map[string]Component),
	}
}

func (cr *ComponentRegistry) RegisterComponent(component Component) {
	cr.Lock()
	defer cr.Unlock()
	cr.components[component.CustomID] = component
}

func (cr *ComponentRegistry) HandleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionMessageComponent {
		cr.RLock()
		defer cr.RUnlock()
		if component, ok := cr.components[i.MessageComponentData().CustomID]; ok {
			component.Run(s, i)
		} else {
			// Component not found
			fmt.Println("Component not found:", i.MessageComponentData().CustomID)
		}
	}
}
