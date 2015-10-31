package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/ovh/tat/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Topic struct
type Topic struct {
	ID               string   `bson:"_id"          json:"_id,omitempty"`
	Topic            string   `bson:"topic"        json:"topic"`
	Description      string   `bson:"description"  json:"description"`
	ROGroups         []string `bson:"roGroups"     json:"roGroups,omitempty"`
	RWGroups         []string `bson:"rwGroups"     json:"rwGroups,omitempty"`
	ROUsers          []string `bson:"roUsers"      json:"roUsers,omitempty"`
	RWUsers          []string `bson:"rwUsers"      json:"rwUsers,omitempty"`
	AdminUsers       []string `bson:"adminUsers"   json:"adminUsers,omitempty"`
	AdminGroups      []string `bson:"adminGroups"  json:"adminGroups,omitempty"`
	History          []string `bson:"history"      json:"history"`
	MaxLength        int      `bson:"maxlength"    json:"maxlength"`
	CanForceDate     bool     `bson:"canForceDate" json:"canForceDate"`
	IsROPublic       bool     `bson:"isROPublic"   json:"isROPublic"`
	DateModification int64    `bson:"dateModification" json:"dateModificationn,omitempty"`
	DateCreation     int64    `bson:"dateCreation" json:"dateCreation,omitempty"`
}

// TopicCriteria struct, used by List Topic
type TopicCriteria struct {
	Skip            int
	Limit           int
	IDTopic         string
	Topic           string
	Description     string
	DateMinCreation string
	DateMaxCreation string
	GetNbMsgUnread  string
}

func buildTopicCriteria(criteria *TopicCriteria, user *User) bson.M {
	var query = []bson.M{}

	if criteria.IDTopic != "" {
		queryIDTopics := bson.M{}
		queryIDTopics["$or"] = []bson.M{}
		for _, val := range strings.Split(criteria.IDTopic, ",") {
			queryIDTopics["$or"] = append(queryIDTopics["$or"].([]bson.M), bson.M{"_id": val})
		}
		query = append(query, queryIDTopics)
	}
	if criteria.Topic != "" {
		queryTopics := bson.M{}
		queryTopics["$or"] = []bson.M{}
		for _, val := range strings.Split(criteria.Topic, ",") {
			queryTopics["$or"] = append(queryTopics["$or"].([]bson.M), bson.M{"topic": val})
		}
		query = append(query, queryTopics)
	}
	if criteria.Description != "" {
		queryDescriptions := bson.M{}
		queryDescriptions["$or"] = []bson.M{}
		for _, val := range strings.Split(criteria.Description, ",") {
			queryDescriptions["$or"] = append(queryDescriptions["$or"].([]bson.M), bson.M{"description": val})
		}
		query = append(query, queryDescriptions)
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

	//groups := user.GetGroups()
	bsonUser := []bson.M{}
	bsonUser = append(bsonUser, bson.M{"roUsers": bson.M{"$in": [1]string{user.Username}}})
	bsonUser = append(bsonUser, bson.M{"rwUsers": bson.M{"$in": [1]string{user.Username}}})
	bsonUser = append(bsonUser, bson.M{"adminUsers": bson.M{"$in": [1]string{user.Username}}})
	userGroups, err := user.GetGroupsOnlyName()
	if err != nil {
		log.Errorf("Error with getting groups for user %s", err)
	} else {
		bsonUser = append(bsonUser, bson.M{"roGroups": bson.M{"$in": userGroups}})
		bsonUser = append(bsonUser, bson.M{"rwGroups": bson.M{"$in": userGroups}})
		bsonUser = append(bsonUser, bson.M{"adminGroups": bson.M{"$in": userGroups}})
	}
	query = append(query, bson.M{"$or": bsonUser})

	if len(query) > 0 {
		return bson.M{"$and": query}
	} else if len(query) == 1 {
		return query[0]
	}
	return bson.M{}
}

func getTopicSelectedFields(isAdmin bool) bson.M {
	if !isAdmin {
		return bson.M{"topic": 1, "description": 1, "isROPublic": 1}
	}
	return bson.M{}
}

// CountTopics return the total number of topics in db
func CountTopics() (int, error) {
	return Store().clTopics.Count()
}

// ListTopics returns list of topics, matching criterias
func ListTopics(criteria *TopicCriteria, user *User) (int, []Topic, error) {
	var topics []Topic

	cursor := listTopicsCursor(criteria, user)
	count, err := cursor.Count()
	if err != nil {
		log.Errorf("Error while count Topics %s", err)
	}

	err = cursor.Select(getTopicSelectedFields(user.IsAdmin)).
		Sort("topic").
		Skip(criteria.Skip).
		Limit(criteria.Limit).
		All(&topics)

	if err != nil {
		log.Errorf("Error while Find Topics %s", err)
	}
	return count, topics, err
}

func listTopicsCursor(criteria *TopicCriteria, user *User) *mgo.Query {
	return Store().clTopics.Find(buildTopicCriteria(criteria, user))
}

// InitPrivateTopic insert topic "/Private"
func InitPrivateTopic() {
	topic := &Topic{
		ID:           bson.NewObjectId().Hex(),
		Topic:        "/Private",
		Description:  "Private Topics",
		DateCreation: time.Now().Unix(),
		MaxLength:    DefaultMessageMaxSize,
		CanForceDate: false,
		IsROPublic:   false,
	}
	err := Store().clTopics.Insert(topic)
	log.Infof("Initialize /Private Topic")
	if err != nil {
		log.Fatalf("Error while initialize /Private Topic %s", err)
	}
}

// Insert creates a new topic. User is read write on topic
func (topic *Topic) Insert(user *User) error {
	err := topic.CheckAndFixName()
	if err != nil {
		return err
	}

	isParentRootTopic, parentTopic, err := topic.getParentTopic()
	if !isParentRootTopic {
		if err != nil {
			return fmt.Errorf("Parent Topic not found %s", topic.Topic)
		}

		// If user create a Topic in /Private, no check or RW to create
		if !strings.HasPrefix(topic.Topic, "/Private/"+user.Username) {
			// check if user can create topic in /topic
			hasRW := parentTopic.IsUserRW(user)
			if !hasRW {
				return fmt.Errorf("No RW access to parent topic %s", parentTopic.Topic)
			}
		}
	} else if !user.IsAdmin { // no parent topic, check admin
		return fmt.Errorf("No write access to create parent topic %s", topic.Topic)
	}

	var existing = &Topic{}
	err = existing.FindByTopic(topic.Topic, true)
	if err == nil {
		return fmt.Errorf("Topic Already Exists : %s", topic.Topic)
	}

	topic.ID = bson.NewObjectId().Hex()
	topic.DateCreation = time.Now().Unix()
	topic.MaxLength = DefaultMessageMaxSize // topic MaxLenth messages
	topic.CanForceDate = false
	topic.IsROPublic = false

	err = Store().clTopics.Insert(topic)
	if err != nil {
		log.Errorf("Error while inserting new topic %s", err)
	}

	h := fmt.Sprintf("create a new topic :%s", topic.Topic)
	err = topic.addToHistory(bson.M{"_id": topic.ID}, user.Username, h)
	if err != nil {
		log.Errorf("Error while inserting history for new topic %s", err)
	}
	err = topic.AddRwUser(user.Username, user.Username, false)
	return err
}

// Get parent topic
// If it is a "root topic", like /myTopic, return true, nil, nil
func (topic *Topic) getParentTopic() (bool, *Topic, error) {
	index := strings.LastIndex(topic.Topic, "/")
	if index == 0 {
		return true, nil, nil
	}
	var nameParent = topic.Topic[0:index]
	var parentTopic = &Topic{}
	err := parentTopic.FindByTopic(nameParent, true)
	if err != nil {
		log.Errorf("Error while fetching parent topic %s", err)
	}
	return false, parentTopic, err
}

// FindByTopic returns topic by topicName.
func (topic *Topic) FindByTopic(topicIn string, isAdmin bool) error {
	topic.Topic = topicIn
	err := topic.CheckAndFixName()
	if err != nil {
		return err
	}
	err = Store().clTopics.Find(bson.M{"topic": topic.Topic}).
		Select(getTopicSelectedFields(isAdmin)).
		One(&topic)
	if err != nil {
		log.Debugf("Error while fetching topic %s", topic.Topic)
	}
	return err
}

// IsTopicExists return true if topic exists, false otherwise
func IsTopicExists(topic string) bool {
	var t = Topic{}
	err := t.FindByTopic(topic, false)
	if err != nil {
		return false // topic does not exist
	}
	return true // topic exists
}

// FindByID return topic, matching given id
func (topic *Topic) FindByID(id string, isAdmin bool) error {
	err := Store().clTopics.Find(bson.M{"_id": id}).
		Select(getTopicSelectedFields(isAdmin)).
		One(&topic)
	if err != nil {
		log.Errorf("Error while fecthing topic with id:%s", id)
	}
	return err
}

// SetParam update param maxLength, canForceDate, isROPublic on topic
func (topic *Topic) SetParam(username string, recursive bool, maxLength int, canForceDate, isROPublic bool) error {

	var selector bson.M

	if recursive {
		selector = bson.M{"topic": bson.RegEx{Pattern: "^" + topic.Topic + ".*$"}}
	} else {
		selector = bson.M{"_id": topic.ID}
	}

	if maxLength <= 0 {
		maxLength = DefaultMessageMaxSize
	}

	_, err := Store().clTopics.UpdateAll(
		selector,
		bson.M{
			"$set": bson.M{
				"maxlength":    maxLength,
				"canForceDate": canForceDate,
				"isROPublic":   isROPublic,
			},
		},
	)

	if err != nil {
		return err
	}
	h := fmt.Sprintf("update param to maxlength:%d, canForceDate:%t, isROPublic:%t", maxLength, canForceDate, isROPublic)
	return topic.addToHistory(selector, username, h)
}

func (topic *Topic) actionOnSet(operand, set, username, admin string, recursive bool, history string) error {

	var selector bson.M

	if recursive {
		selector = bson.M{"topic": bson.RegEx{Pattern: "^" + topic.Topic + ".*$"}}
	} else {
		selector = bson.M{"_id": topic.ID}
	}

	_, err := Store().clTopics.UpdateAll(
		selector,
		bson.M{operand: bson.M{set: username}},
	)

	if err != nil {
		return err
	}
	return topic.addToHistory(selector, admin, history+" "+username)
}

// AddRoUser add a read only user to topic
func (topic *Topic) AddRoUser(admin string, username string, recursive bool) error {
	return topic.actionOnSet("$addToSet", "roUsers", username, admin, recursive, "add to ro")
}

// AddRwUser add a read write user to topic
func (topic *Topic) AddRwUser(admin string, username string, recursive bool) error {
	return topic.actionOnSet("$addToSet", "rwUsers", username, admin, recursive, "add to ro")
}

// AddAdminUser add a read write user to topic
func (topic *Topic) AddAdminUser(admin string, username string, recursive bool) error {
	return topic.actionOnSet("$addToSet", "adminUsers", username, admin, recursive, "add to admin")
}

// RemoveRoUser removes a read only user from topic
func (topic *Topic) RemoveRoUser(admin string, username string, recursive bool) error {
	return topic.actionOnSet("$pull", "roUsers", username, admin, recursive, "remove from ro")
}

// RemoveAdminUser removes a read only user from topic
func (topic *Topic) RemoveAdminUser(admin string, username string, recursive bool) error {
	return topic.actionOnSet("$pull", "roUsers", username, admin, recursive, "remove from admin")
}

// RemoveRwUser removes a read write user from topic
func (topic *Topic) RemoveRwUser(admin string, username string, recursive bool) error {
	return topic.actionOnSet("$pull", "rwUsers", username, admin, recursive, "remove from rw")
}

// AddRoGroup add a read only group to topic
func (topic *Topic) AddRoGroup(admin string, username string, recursive bool) error {
	return topic.actionOnSet("$addToSet", "roGroups", username, admin, recursive, "add to ro")
}

// AddRwGroup add a read write group to topic
func (topic *Topic) AddRwGroup(admin string, username string, recursive bool) error {
	return topic.actionOnSet("$addToSet", "rwGroups", username, admin, recursive, "add to ro")
}

// AddAdminGroup add a admin group to topic
func (topic *Topic) AddAdminGroup(admin string, username string, recursive bool) error {
	return topic.actionOnSet("$addToSet", "adminGroups", username, admin, recursive, "add to admin")
}

// RemoveAdminGroup removes a read write group from topic
func (topic *Topic) RemoveAdminGroup(admin string, username string, recursive bool) error {
	return topic.actionOnSet("$pull", "adminGroups", username, admin, recursive, "remove from admin")
}

// RemoveRoGroup removes a read only group from topic
func (topic *Topic) RemoveRoGroup(admin string, username string, recursive bool) error {
	return topic.actionOnSet("$pull", "roGroups", username, admin, recursive, "remove from ro")
}

// RemoveRwGroup removes a read write group from topic
func (topic *Topic) RemoveRwGroup(admin string, username string, recursive bool) error {
	return topic.actionOnSet("$pull", "rwGroups", username, admin, recursive, "remove from rw")
}

func (topic *Topic) addToHistory(selector bson.M, user string, historyToAdd string) error {
	toAdd := strconv.FormatInt(time.Now().Unix(), 10) + " " + user + " " + historyToAdd
	_, err := Store().clTopics.UpdateAll(
		selector,
		bson.M{"$addToSet": bson.M{"history": toAdd}},
	)
	return err
}

// IsUserRW return true if user can write on a this topic
// Check personal access to topic, and group access
func (topic *Topic) IsUserRW(user *User) bool {
	if utils.ArrayContains(topic.RWUsers, user.Username) ||
		utils.ArrayContains(topic.AdminUsers, user.Username) {
		return true
	}
	userGroups, err := user.GetGroups()
	if err != nil {
		log.Errorf("Error while fetching user groups")
		return false
	}

	var groups []string
	for _, g := range userGroups {
		groups = append(groups, g.Name)
	}

	if utils.ItemInBothArrays(topic.RWGroups, groups) {
		return true
	}

	return false
}

// IsUserReadAccess  return true if user has read access to topic
func (topic *Topic) IsUserReadAccess(user User) bool {
	currentTopic := topic

	if topic.IsROPublic {
		return true
	}

	// if user not admin, reload topic with admin rights
	if !user.IsAdmin {
		currentTopic = &Topic{}
		e := currentTopic.FindByID(topic.ID, true)
		if e != nil {
			return false
		}
	}

	if utils.ArrayContains(currentTopic.ROUsers, user.Username) ||
		utils.ArrayContains(currentTopic.RWUsers, user.Username) ||
		utils.ArrayContains(currentTopic.AdminUsers, user.Username) {
		return true
	}
	userGroups, err := user.GetGroups()
	if err != nil {
		log.Errorf("Error while fetching user groups for user %s", user.Username)
		return false
	}

	var groups []string
	for _, g := range userGroups {
		groups = append(groups, g.Name)
	}

	if utils.ItemInBothArrays(currentTopic.RWGroups, groups) ||
		utils.ItemInBothArrays(currentTopic.ROGroups, groups) ||
		utils.ItemInBothArrays(currentTopic.AdminGroups, groups) {
		return true
	}

	return false
}

// IsUserAdmin return true if user is admin on this topic
// Check personal access to topic, and group access
func (topic *Topic) IsUserAdmin(user *User) bool {
	if utils.ArrayContains(topic.AdminUsers, user.Username) {
		return true
	}

	userGroups, err := user.GetGroups()
	if err != nil {
		log.Errorf("Error while fetching user groups")
		return false
	}

	var groups []string
	for _, g := range userGroups {
		groups = append(groups, g.Name)
	}

	if utils.ItemInBothArrays(topic.AdminGroups, groups) {
		return true
	}

	return false
}

// CheckAndFixNameTopic Add a / to topic name is it is not present
// return an error if length of name is < 4 or > 100
func CheckAndFixNameTopic(topicName string) (string, error) {
	name := strings.TrimSpace(topicName)

	if len(name) < 4 {
		return topicName, fmt.Errorf("Invalid topic lenght (3 or more characters): %s", topicName)
	}

	if len(name) > 1 && string(name[0]) != "/" {
		name = "/" + name
	}

	if len(name)-1 == strings.LastIndex(name, "/") {
		name = name[0 : len(name)-1]
	}

	if len(name) > 100 {
		return topicName, fmt.Errorf("Invalid topic lenght (max 100 characters):%s", topicName)
	}

	return name, nil
}

// CheckAndFixName Add a / to topic name is it is not present
// return an error if length of name is < 4 or > 100
func (topic *Topic) CheckAndFixName() error {
	name, err := CheckAndFixNameTopic(topic.Topic)
	if err != nil {
		return err
	}
	topic.Topic = name
	return nil
}

func changeUsernameOnTopics(oldUsername, newUsername string) {
	changeUsernameOnSet("roUsers", oldUsername, newUsername)
	changeUsernameOnSet("rwUsers", oldUsername, newUsername)
	changeUsernameOnSet("adminUsers", oldUsername, newUsername)
	changeUsernameOnPrivateTopics(oldUsername, newUsername)
}

func changeUsernameOnSet(set, oldUsername, newUsername string) {
	_, err := Store().clTopics.UpdateAll(
		bson.M{set: oldUsername},
		bson.M{"$set": bson.M{set + ".$": newUsername}})

	if err != nil {
		log.Errorf("Error while changes username from %s to %s on Topics (%s) %s", oldUsername, newUsername, set, err)
	}
}

func changeUsernameOnPrivateTopics(oldUsername, newUsername string) error {
	var topics []Topic

	err := Store().clTopics.Find(
		bson.M{
			"topic": bson.RegEx{
				Pattern: "^/Private/" + oldUsername + ".*$", Options: "i",
			}}).All(&topics)

	if err != nil {
		log.Errorf("Error while getting topic with username %s for rename to %s on Topics %s", oldUsername, newUsername, err)
	}

	for _, topic := range topics {
		newTopicName := strings.Replace(topic.Topic, oldUsername, newUsername, 1)
		err := Store().clTopics.Update(
			bson.M{"_id": topic.ID},
			bson.M{"$set": bson.M{"topic": newTopicName}},
		)
		if err != nil {
			log.Errorf("Error while update Topic name from %s to %s :%s", topic.Topic, newTopicName, err)
		}
	}

	return err
}
