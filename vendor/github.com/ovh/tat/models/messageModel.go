package models

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/mvdan/xurls"
	"github.com/ovh/tat/utils"
	"github.com/yesnault/hashtag"
	"gopkg.in/mgo.v2/bson"
)

// DefaultMessageMaxSize is max size of message, can be overrided by topic
var DefaultMessageMaxSize = 140

// Author struct
type Author struct {
	Username string `bson:"username" json:"username"`
	Fullname string `bson:"fullname" json:"fullname"`
}

// Label struct
type Label struct {
	Text  string `bson:"text" json:"text"`
	Color string `bson:"color" json:"color"`
}

// Message struc
type Message struct {
	ID              string    `bson:"_id"             json:"_id"`
	Text            string    `bson:"text"            json:"text"`
	Topics          []string  `bson:"topics"          json:"topics"`
	InReplyOfID     string    `bson:"inReplyOfID"     json:"inReplyOfID"`
	InReplyOfIDRoot string    `bson:"inReplyOfIDRoot" json:"inReplyOfIDRoot"`
	NbLikes         int64     `bson:"nbLikes"         json:"nbLikes"`
	Labels          []Label   `bson:"labels"          json:"labels,omitempty"`
	Likers          []string  `bson:"likers"          json:"likers,omitempty"`
	UserMentions    []string  `bson:"userMentions"    json:"userMentions,omitempty"`
	Urls            []string  `bson:"urls"            json:"urls,omitempty"`
	Tags            []string  `bson:"tags"            json:"tags,omitempty"`
	DateCreation    int64     `bson:"dateCreation"    json:"dateCreation"`
	DateUpdate      int64     `bson:"dateUpdate"      json:"dateUpdate"`
	Author          Author    `bson:"author"          json:"author"`
	Replies         []Message `bson:"-"               json:"replies,omitempty"`
}

// MessageCriteria are used to list messages
type MessageCriteria struct {
	Skip            int
	Limit           int
	TreeView        string
	IDMessage       string
	InReplyOfID     string
	InReplyOfIDRoot string
	AllIDMessage    string // search in IDMessage OR InReplyOfID OR InReplyOfIDRoot
	Text            string
	Topic           string
	Label           string
	NotLabel        string
	AndLabel        string
	Tag             string
	NotTag          string
	AndTag          string
	DateMinCreation string
	DateMaxCreation string
	DateMinUpdate   string
	DateMaxUpdate   string
}

func buildMessageCriteria(criteria *MessageCriteria) bson.M {
	var query = []bson.M{}

	if criteria.IDMessage != "" {
		queryIDMessages := bson.M{}
		queryIDMessages["$or"] = []bson.M{}
		for _, val := range strings.Split(criteria.IDMessage, ",") {
			queryIDMessages["$or"] = append(queryIDMessages["$or"].([]bson.M), bson.M{"_id": val})
		}
		query = append(query, queryIDMessages)
	}
	if criteria.InReplyOfID != "" {
		queryIDMessages := bson.M{}
		queryIDMessages["$or"] = []bson.M{}
		for _, val := range strings.Split(criteria.InReplyOfID, ",") {
			queryIDMessages["$or"] = append(queryIDMessages["$or"].([]bson.M), bson.M{"inReplyOfID": val})
		}
		query = append(query, queryIDMessages)
	}
	if criteria.InReplyOfIDRoot != "" {
		queryIDMessages := bson.M{}
		queryIDMessages["$or"] = []bson.M{}
		for _, val := range strings.Split(criteria.InReplyOfIDRoot, ",") {
			queryIDMessages["$or"] = append(queryIDMessages["$or"].([]bson.M), bson.M{"inReplyOfIDRoot": val})
		}
		query = append(query, queryIDMessages)
	}

	if criteria.AllIDMessage != "" {
		queryIDMessages := bson.M{}
		queryIDMessages["$or"] = []bson.M{}
		for _, val := range strings.Split(criteria.AllIDMessage, ",") {
			queryIDMessages["$or"] = append(queryIDMessages["$or"].([]bson.M), bson.M{"_id": val})
			queryIDMessages["$or"] = append(queryIDMessages["$or"].([]bson.M), bson.M{"inReplyOfID": val})
			queryIDMessages["$or"] = append(queryIDMessages["$or"].([]bson.M), bson.M{"inReplyOfIDRoot": val})
		}
		query = append(query, queryIDMessages)
	}

	if criteria.Text != "" {
		queryTexts := bson.M{}
		queryTexts["$or"] = []bson.M{}
		for _, val := range strings.Split(criteria.Text, ",") {
			queryTexts["$or"] = append(queryTexts["$or"].([]bson.M), bson.M{"text": bson.RegEx{Pattern: "^.*" + regexp.QuoteMeta(val) + ".*$", Options: "im"}})
		}
		query = append(query, queryTexts)
	}
	if criteria.Topic != "" {
		queryTopics := bson.M{}
		queryTopics["$or"] = []bson.M{}
		queryTopics["$or"] = append(queryTopics["$or"].([]bson.M), bson.M{"topics": bson.M{"$in": strings.Split(criteria.Topic, ",")}})
		query = append(query, queryTopics)
	}
	if criteria.Label != "" {
		queryLabels := bson.M{"labels": bson.M{"$elemMatch": bson.M{"text": bson.M{"$in": strings.Split(criteria.Label, ",")}}}}
		query = append(query, queryLabels)
	}
	if criteria.AndLabel != "" {
		queryLabels := bson.M{"labels": bson.M{"$elemMatch": bson.M{"text": bson.M{"$all": strings.Split(criteria.AndLabel, ",")}}}}
		query = append(query, queryLabels)
	}
	if criteria.NotLabel != "" {
		queryLabels := bson.M{}
		queryLabels["$or"] = []bson.M{}
		for _, val := range strings.Split(criteria.NotLabel, ",") {
			queryLabels["$or"] = append(queryLabels["$or"].([]bson.M), bson.M{"labels.text": bson.M{"$ne": val}})
		}
		query = append(query, queryLabels)
	}
	if criteria.Tag != "" {
		queryTags := bson.M{"tags": bson.M{"$in": strings.Split(criteria.Tag, ",")}}
		query = append(query, queryTags)
	}
	if criteria.AndTag != "" {
		queryTags := bson.M{"tags": bson.M{"$all": strings.Split(criteria.AndTag, ",")}}
		query = append(query, queryTags)
	}
	if criteria.NotTag != "" {
		queryTags := bson.M{"tags": bson.M{"$nin": strings.Split(criteria.NotTag, ",")}}
		query = append(query, queryTags)
	}

	var bsonDate = bson.M{}
	if criteria.DateMinCreation != "" {
		i, err := strconv.ParseInt(criteria.DateMinCreation, 10, 64)
		if err != nil {
			log.Errorf("Error while parsing dateMinCreation %s", err)
		}
		tm := time.Unix(i, 0)

		if err == nil {
			bsonDate["$gte"] = tm.Unix()
		} else {
			log.Errorf("Error while parsing dateMinCreation %s", err)
		}
	}
	if criteria.DateMaxCreation != "" {
		i, err := strconv.ParseInt(criteria.DateMaxCreation, 10, 64)
		if err != nil {
			log.Errorf("Error while parsing dateMaxCreation %s", err)
		}
		tm := time.Unix(i, 0)

		if err == nil {
			bsonDate["$lte"] = tm.Unix()
		} else {
			log.Errorf("Error while parsing dateMaxCreation %s", err)
		}
	}
	if len(bsonDate) > 0 {
		query = append(query, bson.M{"dateCreation": bsonDate})
	}

	var bsonDateUpdate = bson.M{}
	if criteria.DateMinUpdate != "" {
		i, err := strconv.ParseInt(criteria.DateMinUpdate, 10, 64)
		if err != nil {
			log.Errorf("Error while parsing dateMinUpdate %s", err)
		}
		tm := time.Unix(i, 0)

		if err == nil {
			bsonDateUpdate["$gte"] = tm.Unix()
		} else {
			log.Errorf("Error while parsing dateMinUpdate %s", err)
		}
	}
	if criteria.DateMaxUpdate != "" {
		i, err := strconv.ParseInt(criteria.DateMaxUpdate, 10, 64)
		if err != nil {
			log.Errorf("Error while parsing dateMaxUpdate %s", err)
		}
		tm := time.Unix(i, 0)

		if err == nil {
			bsonDateUpdate["$lte"] = tm.Unix()
		} else {
			log.Errorf("Error while parsing dateMaxUpdate %s", err)
		}
	}
	if len(bsonDateUpdate) > 0 {
		query = append(query, bson.M{"dateUpdate": bsonDateUpdate})
	}

	if len(query) > 0 {
		return bson.M{"$and": query}
	} else if len(query) == 1 {
		return query[0]
	}
	return bson.M{}
}

// FindByID returns message by given ID
func (message *Message) FindByID(id string) error {
	err := Store().clMessages.Find(bson.M{"_id": id}).One(&message)
	if err != nil {
		log.Errorf("Error while fecthing message with id %s", id)
	}
	return err
}

// ListMessages list messages with given criteria
func ListMessages(criteria *MessageCriteria) ([]Message, error) {
	var messages []Message

	err := Store().clMessages.Find(buildMessageCriteria(criteria)).
		Sort("-dateCreation").
		Skip(criteria.Skip).
		Limit(criteria.Limit).
		All(&messages)

	if err != nil {
		log.Errorf("Error while Find All Messages %s", err)
	}
	if criteria.TreeView == "onetree" {
		return oneTreeMessages(messages, 1)
	} else if criteria.TreeView == "fulltree" {
		return fullTreeMessages(messages, 1)
	}

	return messages, err
}

func oneTreeMessages(messages []Message, nloop int) ([]Message, error) {
	var tree []Message
	if nloop > 25 {
		e := "Infinite loop detected in oneTreeMessages"
		log.Errorf(e)
		return tree, errors.New(e)
	}

	replies := make(map[string][]Message)
	for i := 0; i <= len(messages)-1; i++ {
		if messages[i].InReplyOfIDRoot == "" {
			var replyAdded = false
			for _, msgReply := range replies[messages[i].ID] {
				messages[i].Replies = append(messages[i].Replies, msgReply)
				replyAdded = true
			}
			if replyAdded || nloop > 1 {
				tree = append(tree, messages[i])
				delete(replies, messages[i].ID)
			} else if nloop == 1 && !replyAdded {
				replies[messages[i].ID] = append(replies[messages[i].ID], messages[i])
			}
			continue
		}
		replies[messages[i].InReplyOfIDRoot] = append(replies[messages[i].InReplyOfIDRoot], messages[i])
	}

	if len(replies) == 0 {
		return tree, nil
	}
	t, err := getTree(replies)
	if err != nil {
		return tree, err
	}

	ft, err := oneTreeMessages(t, nloop+1)
	return append(ft, tree...), err
}

func fullTreeMessages(messages []Message, nloop int) ([]Message, error) {
	var tree []Message
	if nloop > 10 {
		e := "Infinite loop detected in fullTreeMessages"
		log.Errorf(e)
		return tree, errors.New(e)
	}

	replies := make(map[string][]Message)
	var alreadyDone []string

	for i := 0; i <= len(messages)-1; i++ {
		if utils.ArrayContains(alreadyDone, messages[i].ID) {
			continue
		}

		var replyAdded = false
		for _, msgReply := range replies[messages[i].ID] {
			messages[i].Replies = append(messages[i].Replies, msgReply)
			delete(replies, messages[i].ID)
			replyAdded = true
		}
		if messages[i].InReplyOfIDRoot == "" {
			if replyAdded || nloop > 1 {
				tree = append(tree, messages[i])
			} else if nloop == 1 && !replyAdded {
				replies[messages[i].ID] = append(replies[messages[i].ID], messages[i])
			}
			continue
		}
		replies[messages[i].InReplyOfID] = append(replies[messages[i].InReplyOfID], messages[i])
		alreadyDone = append(alreadyDone, messages[i].ID)
	}

	if len(replies) == 0 {
		return tree, nil
	}
	t, err := getTree(replies)
	if err != nil {
		return tree, err
	}
	ft, err := fullTreeMessages(t, nloop+1)
	return append(ft, tree...), err
}

func getTree(messagesIn map[string][]Message) ([]Message, error) {
	var messages []Message

	for idMessage := range messagesIn {
		c := &MessageCriteria{
			AllIDMessage: idMessage,
		}
		var msgs []Message
		err := Store().clMessages.Find(buildMessageCriteria(c)).Sort("-dateCreation").All(&msgs)
		if err != nil {
			log.Errorf("Error while Find Messages in getTree %s", err)
			return messages, err
		}
		messages = append(messages, msgs...)
	}

	return messages, nil
}

// Insert a new message on one topic
func (message *Message) Insert(user User, topic Topic, text, inReplyOfID string, dateCreation int64) error {
	message.Text = text
	err := message.CheckAndFixText(topic)
	if err != nil {
		return err
	}
	message.ID = bson.NewObjectId().Hex()
	message.InReplyOfID = inReplyOfID
	dateToStore := time.Now().Unix()

	if dateCreation > 0 {
		if !topic.CanForceDate {
			return fmt.Errorf("You can't force date on topic %s", topic.Topic)
		}

		if !user.IsSystem {
			return fmt.Errorf("You can't force date on topic %s, you're not a system user", topic.Topic)
		}

		if !topic.CanForceDate {
			return fmt.Errorf("Error while converting dateCreation %d - error:%s", dateCreation, err.Error())
		}
		dateToStore = dateCreation
	}

	if inReplyOfID != "" {
		var messageReference = &Message{}
		err = messageReference.FindByID(inReplyOfID)
		if err != nil {
			return err
		}
		if messageReference.InReplyOfID != "" {
			message.InReplyOfIDRoot = messageReference.InReplyOfIDRoot
		} else {
			message.InReplyOfIDRoot = messageReference.ID
		}
		message.Topics = messageReference.Topics

		// if msgRef.dateCreation >= dateToStore -> dateToStore must be after
		if dateToStore <= messageReference.DateCreation {
			dateToStore = messageReference.DateCreation + 1
		}
	} else {
		message.Topics = append(message.Topics, topic.Topic)

		topicDM := "/Private/" + user.Username + "/DM/"
		if strings.HasPrefix(topic.Topic, topicDM) {
			part := strings.Split(topic.Topic, "/")
			if len(part) != 5 {
				log.Errorf("wrong topic name for DM")
				return fmt.Errorf("Wrong topic name for DM:%s", topic.Topic)
			}
			topicInverse := "/Private/" + part[4] + "/DM/" + user.Username
			message.Topics = append(message.Topics, topicInverse)
		}
	}

	message.NbLikes = 0
	var author = Author{}
	author.Username = user.Username
	author.Fullname = user.Fullname
	message.Author = author

	message.DateCreation = dateToStore
	message.DateUpdate = time.Now().Unix()
	message.Tags = hashtag.ExtractHashtags(message.Text)
	message.UserMentions = hashtag.ExtractMentions(message.Text)
	message.Urls = xurls.Strict.FindAllString(message.Text, -1)

	err = Store().clMessages.Insert(message)
	if err != nil {
		log.Errorf("Error while inserting new message %s", err)
	}
	return err
}

// CheckAndFixText truncates to maxLength (parameter on topic) characters
// if len < 1, return error
func (message *Message) CheckAndFixText(topic Topic) error {
	text := strings.TrimSpace(message.Text)
	if len(text) < 1 {
		return fmt.Errorf("Invalid Text:%s", message.Text)
	}

	maxLength := DefaultMessageMaxSize
	if topic.MaxLength > 0 {
		maxLength = topic.MaxLength
	}

	if len(text) > maxLength {
		text = text[0:maxLength]
	}
	message.Text = text
	return nil
}

// Delete deletes a message from database
func (message *Message) Delete() error {
	err := Store().clMessages.Remove(bson.M{"_id": message.ID})

	if err != nil {
		return err
	}
	return nil
}

func (message *Message) getLabel(label string) (int, Label, error) {
	for idx, cur := range message.Labels {
		if cur.Text == label {
			return idx, cur, nil
		}
	}
	l := Label{}
	return -1, l, fmt.Errorf("label %s not found", label)
}

func (message *Message) containsLabel(label string) bool {
	_, _, err := message.getLabel(label)
	if err == nil {
		return true
	}
	return false
}

func (message *Message) getTag(tag string) (int, string, error) {
	for idx, cur := range message.Tags {
		if cur == tag {
			return idx, cur, nil
		}
	}
	return -1, "", fmt.Errorf("tag %s not found", tag)
}

func (message *Message) containsTag(tag string) bool {
	_, _, err := message.getTag(tag)
	if err == nil {
		return true
	}
	return false
}

//AddLabel add a label to a message
//truncated to 20 char in text label
func (message *Message) AddLabel(label string, color string) (Label, error) {
	if len(label) > 20 {
		label = label[0:20]
	}

	if message.containsLabel(label) {
		return Label{}, fmt.Errorf("AddLabel not possible, %s is already a label of this message", label)
	}
	var newLabel = Label{Text: label, Color: color}

	err := Store().clMessages.Update(
		bson.M{"_id": message.ID},
		bson.M{"$set": bson.M{"dateUpdate": time.Now().Unix()}, "$push": bson.M{"labels": newLabel}})

	if err != nil {
		return Label{}, err
	}
	message.Labels = append(message.Labels, newLabel)
	return newLabel, nil
}

// RemoveLabel removes label from on message (label text matching)
func (message *Message) RemoveLabel(label string) error {
	idxLabel, l, err := message.getLabel(label)
	if err != nil {
		return fmt.Errorf("Remove Label is not possible, %s is not a label of this message", label)
	}

	err = Store().clMessages.Update(
		bson.M{"_id": message.ID},
		bson.M{"$set": bson.M{"dateUpdate": time.Now().Unix()}, "$pull": bson.M{"labels": l}})

	if err != nil {
		return err
	}

	message.Labels = append(message.Labels[:idxLabel], message.Labels[idxLabel+1:]...)
	return nil
}

//AddTag add a tag to a message
func (message *Message) AddTag(tag string) error {
	if message.containsTag(tag) {
		return fmt.Errorf("AddTag not possible, %s is already a tag of this message", tag)
	}

	err := Store().clMessages.Update(
		bson.M{"_id": message.ID},
		bson.M{"$set": bson.M{"dateUpdate": time.Now().Unix()}, "$push": bson.M{"tags": tag}})

	if err != nil {
		return err
	}
	message.Tags = append(message.Tags, tag)
	return nil
}

// RemoveTag removes tag from on message
func (message *Message) RemoveTag(tag string) error {
	idxTag, l, err := message.getTag(tag)
	if err != nil {
		return fmt.Errorf("Remove tag is not possible, %s is not a tag of this message", tag)
	}

	err = Store().clMessages.Update(
		bson.M{"_id": message.ID},
		bson.M{"$set": bson.M{"dateUpdate": time.Now().Unix()}, "$pull": bson.M{"tags": l}})

	if err != nil {
		return err
	}

	message.Tags = append(message.Tags[:idxTag], message.Tags[idxTag+1:]...)
	return nil
}

// Like add a like to a message
func (message *Message) Like(user User) error {
	if utils.ArrayContains(message.Likers, user.Username) {
		return fmt.Errorf("Like not possible, %s is already a liker of this message", user.Username)
	}
	err := Store().clMessages.Update(
		bson.M{"_id": message.ID},
		bson.M{"$set": bson.M{"dateUpdate": time.Now().Unix()},
			"$inc":  bson.M{"nbLikes": 1},
			"$push": bson.M{"likers": user.Username}})

	if err != nil {
		return err
	}
	return nil
}

// Unlike removes a like from one message
func (message *Message) Unlike(user User) error {
	if !utils.ArrayContains(message.Likers, user.Username) {
		return fmt.Errorf("Unlike not possible, %s is not a liker of this message", user.Username)
	}
	err := Store().clMessages.Update(
		bson.M{"_id": message.ID},
		bson.M{"$set": bson.M{"dateUpdate": time.Now().Unix()},
			"$inc":  bson.M{"nbLikes": -1},
			"$pull": bson.M{"likers": user.Username}})

	if err != nil {
		return err
	}
	return nil
}

// GetPrivateTopicTaskName return Tasks Topic name of user
func GetPrivateTopicTaskName(user User) string {
	return "/Private/" + user.Username + "/Tasks"
}

func (message *Message) addOrRemoveFromTasks(action string, user User, topic Topic) error {
	if action != "pull" && action != "push" {
		return fmt.Errorf("Wrong action to add or remove tasks:%s", action)
	}
	topicTasksName := GetPrivateTopicTaskName(user)
	idRoot := message.ID
	if message.InReplyOfIDRoot != "" {
		idRoot = message.InReplyOfIDRoot
	}

	_, err := Store().clMessages.UpdateAll(
		bson.M{"$or": []bson.M{bson.M{"_id": idRoot}, bson.M{"inReplyOfIDRoot": idRoot}}},
		bson.M{"$" + action: bson.M{"topics": topicTasksName}})

	if err != nil {
		return err
	}

	msgReply := &Message{}
	text := "Take this thread into my tasks"
	if action == "pull" {
		text = "Remove this thread from my tasks"
	}
	return msgReply.Insert(user, topic, text, idRoot, -1)
}

// AddToTasks add a message to user's tasks Topic
func (message *Message) AddToTasks(user User, topic Topic) error {
	return message.addOrRemoveFromTasks("push", user, topic)
}

// RemoveFromTasks removes a task from user's Tasks Topic
func (message *Message) RemoveFromTasks(user User, topic Topic) error {
	return message.addOrRemoveFromTasks("pull", user, topic)
}

// CountMsgSinceDate return number of messages created on one topic from a given date
func CountMsgSinceDate(topic string, date int64) (int, error) {
	nb, err := Store().clMessages.Find(bson.M{"topics": bson.M{"$in": [1]string{topic}}, "dateCreation": bson.M{"$gte": date}}).Count()
	if err != nil {
		log.Errorf("Error while count message with topic %s and dateCreation lte:%d", topic, date)
	}
	return nb, err
}

func changeUsernameOnMessages(oldUsername, newUsername string) {
	changeAuthorUsernameOnMessages(oldUsername, newUsername)
	changeUsernameOnMessagesTopics(oldUsername, newUsername)
}

func changeAuthorUsernameOnMessages(oldUsername, newUsername string) error {
	_, err := Store().clMessages.UpdateAll(
		bson.M{"author.username": oldUsername},
		bson.M{"$set": bson.M{"author.username": newUsername}})

	if err != nil {
		log.Errorf("Error while update username from %s to %s on Messages %s", oldUsername, newUsername, err)
	}

	return err
}

func changeUsernameOnMessagesTopics(oldUsername, newUsername string) error {
	var messages []Message

	err := Store().clMessages.Find(
		bson.M{
			"topics": bson.RegEx{Pattern: "^/Private/" + oldUsername + "/", Options: "i"},
		}).All(&messages)

	if err != nil {
		log.Errorf("Error while getting messages to update username from %s to %s on Topics %s", oldUsername, newUsername, err)
	}

	for _, msg := range messages {
		msg.Topics = []string{}
		for _, topic := range msg.Topics {
			newTopicName := strings.Replace(topic, oldUsername, newUsername, 1)
			msg.Topics = append(msg.Topics, newTopicName)
		}

		err := Store().clMessages.Update(
			bson.M{"_id": msg.ID},
			bson.M{"$set": bson.M{"topics": msg.Topics}},
		)

		if err != nil {
			log.Errorf("Error while update topic on message %s name from username %s to username %s :%s", msg.ID, oldUsername, newUsername, err)
		}
	}

	return err
}

// CountMessages returns the total number of messages in db
func CountMessages() (int, error) {
	return Store().clMessages.Count()
}

// DistributionMessages returns distribution of messages per topic
func DistributionMessages(col string) ([]bson.M, error) {
	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id": bson.M{col: "$" + col},
				"count": bson.M{
					"$sum": 1,
				},
			},
		},
		{
			"$sort": bson.M{
				"count": -1,
			},
		},
	}
	pipe := Store().clMessages.Pipe(pipeline)
	results := []bson.M{}

	err := pipe.All(&results)

	return results, err
}
