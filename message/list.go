package message

import (
	"fmt"
	"os"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
)

var (
	dateCreation int

	treeView                  string
	criteriaIDMessage         string
	criteriaInReplyOfID       string
	criteriaInReplyOfIDRoot   string
	criteriaAllIDMessage      string
	criteriaText              string
	criteriaTopic             string
	criteriaLabel             string
	criteriaNotLabel          string
	criteriaAndLabel          string
	criteriaTag               string
	criteriaNotTag            string
	criteriaAndTag            string
	criteriaDateMinCreation   string
	criteriaDateMaxCreation   string
	criteriaDateMinUpdate     string
	criteriaDateMaxUpdate     string
	criteriaUsername          string
	criteriaLimitMinNbReplies string
	criteriaLimitMaxNbReplies string
)

func init() {
	cmdMessageList.Flags().StringVarP(&treeView, "treeView", "", "", "Tree View of messages: onetree or fulltree. Default: notree")
	cmdMessageList.Flags().StringVarP(&criteriaIDMessage, "idMessage", "", "", "Search by IDMessage")
	cmdMessageList.Flags().StringVarP(&criteriaInReplyOfID, "inReplyOfID", "", "", "Search by IDMessage InReply")
	cmdMessageList.Flags().StringVarP(&criteriaInReplyOfIDRoot, "inReplyOfIDRoot", "", "", "Search by IDMessage IdRoot")
	cmdMessageList.Flags().StringVarP(&criteriaAllIDMessage, "allIDMessage", "", "", "Search in All ID Message (idMessage, idReply, idRoot)")
	cmdMessageList.Flags().StringVarP(&criteriaText, "text", "", "", "Search by text")
	cmdMessageList.Flags().StringVarP(&criteriaTopic, "topic", "", "", "Search by topic")
	cmdMessageList.Flags().StringVarP(&criteriaLabel, "label", "", "", "Search by label: could be labelA,labelB")
	cmdMessageList.Flags().StringVarP(&criteriaNotLabel, "notLabel", "", "", "Search by label (exclude): could be labelA,labelB")
	cmdMessageList.Flags().StringVarP(&criteriaAndLabel, "andLabel", "", "", "Search by label (and) : could be labelA,labelB")
	cmdMessageList.Flags().StringVarP(&criteriaTag, "tag", "", "", "Search by tag : could be tagA,tagB")
	cmdMessageList.Flags().StringVarP(&criteriaNotTag, "notTag", "", "", "Search by tag (exclude) : could be tagA,tagB")
	cmdMessageList.Flags().StringVarP(&criteriaAndTag, "andTag", "", "", "Search by tag (and) : could be tagA,tagB")
	cmdMessageList.Flags().StringVarP(&criteriaDateMinCreation, "dateMinCreation", "", "", "Search by dateMinCreation (timestamp)")
	cmdMessageList.Flags().StringVarP(&criteriaDateMaxCreation, "dateMaxCreation", "", "", "Search by dateMaxCreation (timestamp)")
	cmdMessageList.Flags().StringVarP(&criteriaDateMinUpdate, "dateMinUpdate", "", "", "Search by dateMinUpdate (timestamp)")
	cmdMessageList.Flags().StringVarP(&criteriaDateMaxUpdate, "dateMaxUpdate", "", "", "Search by dateMaxUpdate (timestamp)")
	cmdMessageList.Flags().StringVarP(&criteriaUsername, "username", "", "", "Search by username : could be usernameA,usernameB")
	cmdMessageList.Flags().StringVarP(&criteriaLimitMinNbReplies, "limitMinNbReplies", "", "", "In onetree mode, filter root messages with more or equals minNbReplies")
	cmdMessageList.Flags().StringVarP(&criteriaLimitMaxNbReplies, "limitMaxNbReplies", "", "", "In onetree mode, filter root messages with min or equals maxNbReplies")
}

var cmdMessageList = &cobra.Command{
	Use:     "list",
	Short:   "List all messages on one topic: tatcli msg list <Topic> <skip> <limit>",
	Aliases: []string{"l"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			skip, limit := internal.GetSkipLimit(args)
			messagesList(args[0], skip, limit)
		} else {
			fmt.Fprintf(os.Stderr, "Invalid argument to list message: See tatcli msg list --help\n")
		}
	},
}

func messagesList(topic string, skip, limit string) {
	c := ""
	if treeView != "" {
		c = c + "&treeView=" + treeView
	}
	if criteriaIDMessage != "" {
		c = c + "&idMessage=" + criteriaIDMessage
	}
	if criteriaInReplyOfID != "" {
		c = c + "&inReplyOfID=" + criteriaInReplyOfID
	}
	if criteriaInReplyOfIDRoot != "" {
		c = c + "&inReplyOfIDRoot=" + criteriaInReplyOfIDRoot
	}
	if criteriaAllIDMessage != "" {
		c = c + "&allIDMessage=" + criteriaAllIDMessage
	}
	if criteriaText != "" {
		c = c + "&text=" + criteriaText
	}
	if criteriaTopic != "" {
		c = c + "&topic=" + criteriaTopic
	}
	if criteriaLabel != "" {
		c = c + "&label=" + criteriaLabel
	}
	if criteriaNotLabel != "" {
		c = c + "&notLabel=" + criteriaNotLabel
	}
	if criteriaAndLabel != "" {
		c = c + "&andLabel=" + criteriaAndLabel
	}
	if criteriaTag != "" {
		c = c + "&tag=" + criteriaTag
	}
	if criteriaNotTag != "" {
		c = c + "&notTag=" + criteriaNotTag
	}
	if criteriaAndTag != "" {
		c = c + "&andTag=" + criteriaAndTag
	}
	if criteriaDateMinCreation != "" {
		c = c + "&dateMinCreation=" + criteriaDateMinCreation
	}
	if criteriaDateMaxCreation != "" {
		c = c + "&dateMaxCreation=" + criteriaDateMaxCreation
	}
	if criteriaDateMinUpdate != "" {
		c = c + "&dateMinUpdate=" + criteriaDateMinUpdate
	}
	if criteriaDateMaxUpdate != "" {
		c = c + "&dateMaxUpdate=" + criteriaDateMaxUpdate
	}
	if criteriaUsername != "" {
		c = c + "&username=" + criteriaUsername
	}
	if criteriaLimitMinNbReplies != "" {
		c = c + "&limitMinNbReplies=" + criteriaLimitMinNbReplies
	}
	if criteriaLimitMaxNbReplies != "" {
		c = c + "&limitMaxNbReplies=" + criteriaLimitMaxNbReplies
	}
	fmt.Print(internal.GetWantReturn(fmt.Sprintf("/messages/%s?skip=%s&limit=%s%s", topic, skip, limit, c)))
}
