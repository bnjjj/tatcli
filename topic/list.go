package topic

import (
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var (
	criteriaTopic           string
	criteriaIDTopic         string
	criteriaDescription     string
	criteriaDateMinCreation string
	criteriaDateMaxCreation string
	criteriaGetNbMsgUnread  string
	criteriaGetForTatAdmin  string
)

func init() {
	cmdTopicList.Flags().StringVarP(&criteriaTopic, "topic", "", "", "Search by Topic name, example: /topicA")
	cmdTopicList.Flags().StringVarP(&criteriaIDTopic, "idTopic", "", "", "Search by id of topic")
	cmdTopicList.Flags().StringVarP(&criteriaDescription, "description", "", "", "Search by description of topic")
	cmdTopicList.Flags().StringVarP(&criteriaDateMinCreation, "dateMinCreation", "", "", "Filter result on dateCreation, timestamp Unix format")
	cmdTopicList.Flags().StringVarP(&criteriaDateMaxCreation, "dateMaxCreation", "", "", "Filter result on dateCreation, timestamp Unix Format")
	cmdTopicList.Flags().StringVarP(&criteriaGetNbMsgUnread, "getNbMsgUnread", "", "", "If true, add new array to return, topicsMsgUnread with topic:nbUnreadMsgSinceLastPresenceOnTopic")
	cmdTopicList.Flags().StringVarP(&criteriaGetForTatAdmin, "getForTatAdmin", "", "", "(AdminOnly) If true, and requester is a Tat Admin, returns all topics (except /Private/*) without checking user / group access (RO or RW on Topic)")
}

var cmdTopicList = &cobra.Command{
	Use:     "list",
	Short:   "List all topics: tatcli topic list [<skip>] [<limit>], tatcli topic list -h for see all criterias",
	Aliases: []string{"l"},
	Run: func(cmd *cobra.Command, args []string) {
		skip := "0"
		limit := "10"
		if len(args) >= 2 {
			skip = args[0]
			limit = args[1]
		}

		topicsList(skip, limit)
	},
}

func topicsList(skip string, limit string) {
	n := ""
	if criteriaTopic != "" {
		n += "&topic=" + criteriaTopic
	}
	if criteriaIDTopic != "" {
		n += "&idTopic=" + criteriaIDTopic
	}
	if criteriaDescription != "" {
		n += "&Description=" + criteriaDescription
	}
	if criteriaDateMinCreation != "" {
		n += "&DateMinCreation=" + criteriaDateMinCreation
	}
	if criteriaDateMaxCreation != "" {
		n += "&DateMaxCreation=" + criteriaDateMaxCreation
	}
	if criteriaGetNbMsgUnread != "" {
		n += "&getNbMsgUnread=" + criteriaGetNbMsgUnread
	}
	if criteriaGetForTatAdmin == "true" {
		n += "&getForTatAdmin=" + criteriaGetForTatAdmin
	}

	fmt.Print(internal.GetWantReturn(fmt.Sprintf("/topics?skip=%s&limit=%s%s", skip, limit, n)))
}
