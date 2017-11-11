package roborooney

import (
	"fmt"

	"github.com/arashout/mlpapi"
)

func (robo *RoboRooney) handlerListCommand() string {
	robo.updateTracker()

	textListSlots := ""
	pitchSlots := robo.tracker.retrieveAll()
	for _, pitchSlot := range pitchSlots {
		textSlot := fmt.Sprintf("%s\n", formatSlotMessage(pitchSlot.pitch, pitchSlot.slot))
		textListSlots += textSlot
	}
	return textListSlots
}

func (robo *RoboRooney) handlerUnseenCommand(fromTicker bool) string {
	robo.updateTracker()

	textListSlots := ""
	pitchSlots := robo.tracker.retrieveUnseen()

	if len(pitchSlots) == 0 && !fromTicker {
		return "No new slots are available"
	}

	for _, pitchSlot := range pitchSlots {
		textSlot := fmt.Sprintf("%s\n", formatSlotMessage(pitchSlot.pitch, pitchSlot.slot))
		textListSlots += textSlot
	}
	return textListSlots
}

func (robo *RoboRooney) handlerCheckoutCommand(msgText string) string {
	robo.updateTracker()

	pitchSlotID := regexPitchSlotID.FindString(msgText)
	if pitchSlotID != "" {
		pitchSlot, err := robo.tracker.retrieve(pitchSlotID)
		if err != nil {
			return "Pitch-Slot ID not found. Try listing all available bookings again"
		}
		return mlpapi.GetSlotCheckoutLink(pitchSlot.pitch, pitchSlot.slot)
	}
	return "No Pitch-Slot ID found in message. Make sure it is formatted correctly."
}

func (robo *RoboRooney) handlerPollCommand() string {
	pitchSlots := robo.tracker.retrieveAll()

	if len(pitchSlots) == 0 {
		return "No slots available for polling\nTry checking availablity first."
	}

	textPoll := "/poll 'Which time(s) works best?' "

	for _, pitchSlot := range pitchSlots {
		optionString := fmt.Sprintf(" \"%s\" ", formatSlotMessage(pitchSlot.pitch, pitchSlot.slot))
		textPoll += optionString
	}

	return textPoll
}

func (robo *RoboRooney) handlerRulesCommand() string {
	textRules := ""
	for _, rule := range robo.rules {
		textRules += "-" + rule.Description + "\n"
	}
	return textRules
}

func (robo *RoboRooney) handlerPitchesCommand() string {
	textPitches := ""
	for _, pitch := range robo.pitches {
		textPitches += "-" + pitch.Name + "\n"
	}
	return textPitches
}
