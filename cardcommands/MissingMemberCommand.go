package cardcommands

import "github.com/adlio/trello"

type MissingMemberCommand struct {
	Members []string
}

func (m MissingMemberCommand) CheckCard(card *trello.Card) bool {
	for _, newMember := range m.Members {
		found := false
		existingMembers := card.IDMembers

		for _, existingMember := range existingMembers {
			found = newMember == existingMember
			if found {
				break
			}
		}

		if !found {
			return true
		}
	}

	return false
}

func (m MissingMemberCommand) UpdateCard(card *trello.Card) error {
	errOut := error(nil)

	for _, newMember := range m.Members {
		found := false

		existingMembers := card.IDMembers
		for _, existingMember := range existingMembers {
			found = newMember == existingMember
			if found {
				break
			}
		}

		if !found {
			_, err := card.AddMemberID(newMember)
			if err != nil {
				errOut = err
			}
		}
	}

	return errOut
}
