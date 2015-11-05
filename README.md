# Description
Tat Command Line Interface.

See Tat Engine for more information: https://github.com/tat/tatcli

# How to build
```
git clone https://github.com/ovh/tatcli.git && cd tatcli
go get && go build && ./tatcli -h
```

# Usage
## Documentation

```
Usage:
  tatcli [command]

Available Commands:
  config      Config commands: tatcli config --help
  group       Group commands: tatcli group --help
  message     Manipulate messages: tatcli message --help
  presence    Presence commands: tatcli presence --help
  socket      Socket commands: tatcli socket --help
  stats       Stats commands (admin only): tatcli stats --help
  topic       Topic commands: tatcli topic --help
  update      Update tatcli to the latest release version: tatcli update
  user        User commands: tatcli user --help
  version     Display Version of tatcli and tat engine if configured : tatcli version
  help        Help about any command

Flags:
  -c, --configFile="$HOME/.tatcli/config.json": configuration file, default is $HOME/.tatcli/config.json
  -h, --help=false: help for tatcli
  -p, --password="": password, facultative if you have a $HOME/.tatcli/config.json file
  -t, --pretty=false: Pretty Print Json Output
  -l, --sslInsecureSkipVerify=false: Skip certificate check with SSL connection
      --url="": URL Tat Engine, facultative if you have a $HOME/.tatcli/config.json file
  -u, --username="": username, facultative if you have a $HOME/.tatcli/config.json file
  -v, --verbose=false: verbose output


Use "tatcli [command] --help" for more information about a command.
```

### Command Config

```
Config commands: tatcli config <command>

Usage:
  tatcli config [command]

Aliases:
  config, c


Available Commands:
  template    Write a template configuration file in $HOME/.tatcli/config.json: tatcli config template
  show        Show Configuration: tatcli config show

Flags:
  -h, --help=false: help for config

Global Flags: see tatcli -h

Use "tatcli config [command] --help" for more information about a command.

```

### Command Group

```
Group commands: tatcli group <command>

Usage:
  tatcli group [command]

Aliases:
  group, g


Available Commands:
  list            List all groups: tatcli group list <skip> <limit>
  create          create a new group: tatlic group create <groupname> <description>
  addUser         Add Users to a group: tacli group addUser <groupname> <username1> [<username2> ... ]
  deleteUser      Delete Users from a group: tacli group deleteUser <groupname> <username1> [<username2> ... ]
  addAdminUser    Add Admin Users to a group: tacli group addAdminUser <groupname> <username1> [<username2> ... ]
  deleteAdminUser Delete Admin Users from a group: tacli group deleteAdminUser <groupname> <username1> [<username2> ... ]

Flags:
  -h, --help=false: help for group

Global Flags: see tatcli -h

Use "tatcli group [command] --help" for more information about a command.

```

### Command Message

```
Manipulate messages: tatcli message <command>

Usage:
  tatcli message [command]

Aliases:
  message, m, msg


Available Commands:
  list        List all messages on one topic: tatcli msg list <Topic> <skip> <limit>
  add         tatcli message add [--dateCreation=timestamp] <topic> <my message>
  reply       Reply to a message: tatcli message reply <topic> <inReplyOfId> <my message...>
  bookmark    Bookmark a message to a topic: tatcli message bookmark /Private/username/bookmarks/sub-topic idMessage
  delete      Remove a message (or bookmark) from Private Topic: tatcli message delete <idMessage>
  update      Update a message (if it's enabled on topic): tatcli message update <topic> <idMessage> <my message...>
  task        Create a task from one message to a topic: tatcli message task /Private/username/tasks/sub-topic idMessage
  untask      Remove a message from tasks: tatcli message untask /Private/username/tasks idMessage
  like        Like a message: tatcli message like <idMessage>
  unlike      Unlike a message: tatcli message unlike <idMessage>
  label       Add a label to a message: tatcli message label <idMessage> <colorInHexa> <my Label>
  unlabel     Remove a label from a message: tatcli message unlabel <idMessage> <my Label>
  tag         Add a tag to a message (user system with rights only): tatcli message tag <idMessage> <my Tag>
  untag       Remove a tag from a message (user system with rights only): tatcli message untag <idMessage> <myTag>
  list        List all messages on one topic: tatcli msg list <Topic> <skip> <limit>
  read        List all messages on one public topic (read only): tatcli msg read <Topic> <skip> <limit>

Flags:
  -h, --help=false: help for message

Global Flags: see tatcli -h

Use "tatcli message [command] --help" for more information about a command.

```

#### Command Message list

```
List all messages of a topic: tatcli msg list <Topic> <skip> <limit>

Usage:
  tatcli message list [flags]

Aliases:
  list, l


Flags:
      --allIDMessage="": Search in All ID Message (idMessage, idReply, idRoot)
      --andLabel="": Search by label (and) : could be labelA,labelB
      --andTag="": Search by tag (and) : could be tagA,tagB
      --dateMaxCreation="": Search by dateMaxCreation (timestamp)
      --dateMaxUpdate="": Search by dateMaxUpdate (timestamp)
      --dateMinCreation="": Search by dateMinCreation (timestamp)
      --dateMinUpdate="": Search by dateMinUpdate (timestamp)
  -h, --help=false: help for list
      --idMessage="": Search by IDMessage
      --inReplyOfID="": Search by IDMessage InReply
      --inReplyOfIDRoot="": Search by IDMessage IdRoot
      --label="": Search by label: could be labelA,labelB
      --notLabel="": Search by label (exclude): could be labelA,labelB
      --notTag="": Search by tag (exclude) : could be tagA,tagB
      --tag="": Search by tag : could be tagA,tagB
      --text="": Search by text
      --topic="": Search by topic
      --treeView="": Tree View of messages: onetree or fulltree. Default: notree
      --username="": Search by username : could be usernameA,usernameB

```

### Command Presence

```
Presence commands: tatcli presence [<command>]

Usage:
  tatcli presence [command]

Aliases:
  presence, p


Available Commands:
  add         Add a new presence on one topic with status (online, offline, busy): tatcli presence add <topic> <status>
  list        List all presences on one topic: tatcli presence list <topic> [<skip>] [<limit>]

Flags:
  -h, --help=false: help for presence

Global Flags: see tatcli -h

```

### Command Socket

```
Socket commands: tatcli socket [<command>]

Usage:
  tatcli socket [command]

Aliases:
  socket, s


Available Commands:
  dump        Dump websocket admin variables: tatcli socket dump
  messages    Open websocket and get events messages on one or many topics: tatcli socket messages <<topic>|all|onetree|fulltree>> [topic]...
  messagesNew Open websocket and get events on new messages on one or many topics: tatcli socket messagesNew <<topic>|all>> [topic]...
  interactive Interactive mode Websocket: tatcli socket interactive (or tatcli socket i)
  users       Open websocket and get events users (admin only): tatcli socket users

Flags:
  -h, --help=false: help for socket

Global Flags: see tatcli -h

```

### Command Topic

```
Topic commands: tatcli topic [command]

Usage:
  tatcli topic [command]

Aliases:
  topic, t


Available Commands:
  list             List all topics: tatcli topic list [<skip>] [<limit>] [<true>] if true, return unread cound msg
  create           Create a new topic: tatcli create <topic> <description of topic>
  delete           Delete a new topic: tatcli delete <topic>
  addRoUser        Add Read Only Users to a topic: tatcli topic addRoUser [--recursive] <topic> <username1> [username2]...
  addRwUser        Add Read Write Users to a topic: tatcli topic addRwUser [--recursive] <topic> <username1> [username2]...
  addAdminUser     Add Admin Users to a topic: tatcli topic addAdminUser [--recursive] <topic> <username1> [username2]...
  deleteRoUser     Delete Read Only Users from a topic: tatcli topic deleteRoUser [--recursive] <topic> <username1> [username2]...
  deleteRwUser     Delete Read Write Users from a topic: tatcli topic deleteRwUser [--recursive] <topic> <username1> [username2]...
  deleteAdminUser  Delete Admin Users from a topic: tatcli topic deleteAdminUser [--recursive] <topic> <username1> [username2]...
  addRoGroup       Add Read Only Groups to a topic: tatcli topic addRoGroup [--recursive] <topic> <groupname1> [<groupname2>]...
  addRwGroup       Add Read Write Groups to a topic: tatcli topic addRwGroup [--recursive] <topic> <groupname1> [<groupname2>]...
  addAdminGroup    Add Admin Groups to a topic: tatcli topic addAdminGroup [--recursive] <topic> <groupname1> [groupname2]...
  deleteRoGroup    Delete Read Only Groups from a topic: tatcli topic deleteRoGroup [--recursive] <topic> <groupname1> [<groupname2>]...
  deleteRwGroup    Delete Read Write Groups from a topic: tatcli topic deleteRwGroup [--recursive] <topic> <groupname1> [<groupname2>]...
  deleteAdminGroup Delete Admin Groups from a topic: tatcli topic deleteAdminGroup [--recursive] <topic> <groupname1> [<groupname2>]...
  addParameter     Add Parameter to a topic: tatcli topic addParameter [--recursive] <topic> <key>:<value> [<key2>:<value2>]...
  deleteParameter  Remove Parameter to a topic: tatcli topic deleteParameter [--recursive] <topic> <key> [<key2>]...
  parameter        Update param on one topic: tatcli topic param [--recursive] <topic> <maxLength> <canForceDate> <canUpdateMsg> <canDeleteMsg> <canUpdateAllMsg> <canDeleteAllMsg> <isROPublic>

Flags:
  -h, --help=false: help for topic

Global Flags: see tatcli -h

```

### Command Stats

```
Stats commands (admin only): tatcli stats [<command>]

Usage:
  tatcli stats [command]

Aliases:
  stats, stat


Available Commands:
  count              Count all messages, groups, presences, users, groups, topics: tatcli stats count
  distribution       Distribution of messages per topics: tatcli stats distribution
  dbstats            DB Stats: tatcli stats dbstats
  dbServerStatus     DB Stats: tatcli stats dbServerStatus
  dbReplSetGetConfig DB Stats: tatcli stats dbReplSetGetConfig
  dbReplSetGetStatus DB Stats: tatcli stats dbReplSetGetStatus
  dbCollections      DB Stats on each collection: tatcli stats dbCollections
  dbSlowestQueries   DB Stats slowest Queries: tatcli stats dbSlowestQueries
  instance           Info about current instance of engine

Flags:
  -h, --help=false: help for stats

Global Flags: see tatcli -h

```

### Command Update

```
tatcli update

Usage:
  tatcli update [flags]
  tatcli update [command]

Aliases:
  update, up

Available Commands:
  snapshot    Update tatcli to latest snapshot version: tatcli update snapshot

Flags:
  -h, --help=false: help for update

Global Flags: see tatcli -h

```

### Command User

```
User commands: tatcli user <command>

Usage:
  tatcli user [command]

Aliases:
  user, u


Available Commands:
  list                      List all users: tatcli user list [<skip>] [<limit>]
  me                        Get Information about you: tatcli user me
  contacts                  Get contacts presences since n seconds: tatcli user contacts <seconds>
  addContact                Add a contact: tatcli user addContact <contactUsername>
  removeContact             Remove a contact: tatcli user removeContact <contactUsername>
  addFavoriteTopic          Add a favorite Topic: tatcli user addFavoriteTopic <topicName>
  removeFavoriteTopic       Remove a favorite Topic: tatcli user removeFavoriteTopic <topicName>
  enableNotificationsTopic  Enable notifications on a topic: tatcli user enableNotificationsTopic <topicName>
  disableNotificationsTopic Disable notifications on a topic: tatcli user disableNotificationsTopic <topicName>
  addFavoriteTag            Add a favorite Tag: tatcli user addFavoriteTag <tag>
  removeFavoriteTag         Remove a favorite Tag: tatcli user removeFavoriteTag <tag>
  add                       Add a user: tatcli user add <username> <email> <fullname>
  reset                     Ask for Reset a password: tatcli user reset <username> <email>
  resetSystemUser           Reset password for a system user (admin only): tatcli user resetSystemUser <username>
  convert                   Convert a user to a system user (admin only): tatcli user convert <username> <canWriteNotifications>
  archive                   Archive a user (admin only): tatcli user archive <username>
  rename                    Rename username of a user (admin only): tatcli user rename <oldUsername> <newUsername>
  update                    Update Fullname and Email of a user (admin only): tatcli user update <username> <newEmail> <newFullname>
  setAdmin                  Grant user to Tat admin (admin only): tatcli user setAdmin <username>
  verify                    Verify account: tatcli user verify [--save] <username> <tokenVerify>
  check                     Check Private Topics and Default Group on one user (admin only): tatcli user check <username> <fixPrivateTopics> <fixDefaultGroup>

Flags:
  -h, --help=false: help for user

Global Flags: see tatcli -h

```

### Command Version

```
tatcli version

Usage:
  tatcli version [flags]

Aliases:
  version, v


Flags:
  -h, --help=false: help for version
      --versionNewLine=true: New line after version number. If true, display Version Engine too

Global Flags: see tatcli -h

```

## Examples
### Credentials
Config file is under $HOME/.tatcli/config.json
You can create it with this command:
```
tatcli config template
```

Template is:
```
{
  "username":"myUsername",
  "password":"myPassword",
  "url":"http://urltat:port"
}
```

### Message

#### Create a message
```
tatcli message add /topic my message
```

If you are a `system user`, you can force date creation. Date as timestamp
```
tatcli message add --dateCreation=11111 /topic my message
```

#### Reply to a message
```
tatcli message reply idOfMessage my message
```

#### Like a message
```
tatcli message like idOfMessage
```

#### Unlike a message
```
tatcli message unlike idOfMessage
```

#### Add a label to a message
```
tatcli message label idOfMessage color myLabel
```

#### Remove a label from a message
```
tatcli message unlabel idOfMessage myLabel
```

#### Add a tag to a message

Only for `system user` on his messages

```
tatcli message tag idOfMessage myTag
```

#### Remove a tag from a message

Only for `system user` on his messages

```
tatcli message untag idOfMessage myTag
```

#### Bookmark a message
```
tatcli message bookmark /Private/username/Bookmarks/subtopic idOfMessage
```

#### Unbookmark a message
```
tatcli message unbookmark idOfMessage
```

#### Create a task from one message
```
tatcli message task /Private/username/Tasks idOfMessage
```

#### Remove a message from tasks
```
tatcli message untask /Private/username/Tasks idOfMessage

```


#### Getting message
```
tatcli message list /topic
tatcli message list /topic 0 10
```

#### Getting message on one public topic (access read only)

```
tatcli message read /topic
tatcli message read /topic 0 10
```

### User
#### Create a user
```
tatcli user add username email fullname
```

#### Verify account
```
tatcli user verify username tokenVerify
```

For saving configuration in $HOME/.tatcli/config.json file
```
tatcli user verify --save username tokenVerify
```

#### Ask for reset password
```
tatcli user reset username email
```

#### Get information about me
```
tatcli user me
```

#### Get contacts presences since n seconds: tatcli user contacts <seconds>
```
tatcli user contacts 15
```

#### Add a favorite tag
```
tatcli user addFavoriteTag myTag
```

#### Remove a favorite tag
```
tatcli user removeFavoriteTag myTag
```

#### Add a favorite topic
```
tatcli user addFavoriteTopic /topic/sub-topic
```

#### Remove a favorite topic
```
tatcli user removeFavoriteTopic /topic/sub-topic
```

#### Enable notifications on a topic

Notifications are by default enabled on topic

```
tatcli user enableNotificationsTopic /topic/sub-topic
```

#### Disable notifications on a topic
```
tatcli user disableNotificationsTopic /topic/sub-topic
```

#### List Users
```
tatcli user list
```

with groups (admin only):

```
tatcli user list --withGroups
```

#### Convert to a system user (Admin only)
```
tatcli user convert usernameToConvertSystem flagCanWriteOnNotificationsTopics
```
flagCanWriteOnNotificationsTopics could be true or false

#### Grant a user to Tat Admin (Admin only)
```
tatcli user setAdmin usernameToGrant
```

#### Archive a user (Admin only)
```
tatcli user archive usernameToArchive
```

#### Rename a username  (Admin only)
```
tatcli user rename oldUsername newUsername
```

#### Update fullname and email (Admin only)
```
tatcli user update username newEmail newFirstname newLastname
```

#### Check a user (Admin only)

Check Private Topics and Default Group on one user:

```
tatcli user check <username> <fixPrivateTopics> <fixDefaultGroup>
```

Example :

```
tatcli check username true true
```

### Group

#### Create a group
```
tatcli group add groupname description of group
```

#### Add user to a group
```
tatcli group addUser groupname username
```

#### Delete a user from a group
```
tatcli group deleteUser groupname username
```

### Topic
#### Create a Topic
```
tatcli topic add /topic topic description
```

#### Delete a Topic
```
tatcli topic delete /topic
```

#### Getting Topics List
```
tatcli topic list
tatcli topic list skip limit
tatcli topic list skip limit true
```
if true, return nb unread messages

#### Add a read only user to a topic
```
tatcli topic addRoUser /topic username
tatcli topic addRoUser /topic username1 username2
```
#### Add a read write user to a topic
```
tatcli topic addRwUser /topic username
tatcli topic addRwUser /topic username1 username2
```

#### Add an admin user to a topic
```
tatcli topic addAdminUser /topic username
tatcli topic addAdminUser /topic username1 username2
```


#### Delete a read only user from a topic
```
tatcli topic deleteRoUser /topic username
tatcli topic deleteRoUser /topic username1 username2
```

#### Delete a read write user from a topic
```
tatcli topic deleteRwUser /topic username
tatcli topic deleteRwUser /topic username1 username2
```

#### Delete an admin user from a topic
```
tatcli topic deleteAdminUser /topic username
tatcli topic deleteAdminUser /topic username1 username2
```

#### Add a read only group to a topic
```
tatcli topic addRoGroup /topic groupname
tatcli topic addRoGroup /topic groupname1 groupname2
```

#### Add a read write group to a topic
```
tatcli topic addRwGroup /topic groupname
tatcli topic addRwGroup /topic groupname1 groupname2
```

#### Add an admin group to a topic
```
tatcli topic addAdminGroup /topic groupname
tatcli topic addAdminGroup /topic groupname1 groupname2
```

#### Delete a read only group from a topic
```
tatcli topic deleteRoGroup /topic groupname
tatcli topic deleteRoGroup /topic groupname1 groupname2
```

#### Delete a read write group from a topic
```
tatcli topic deleteRwGroup /topic groupname
tatcli topic deleteRwGroup /topic groupname1 groupname2
```

#### Delete an admin group from a topic
```
tatcli topic deleteAdminGroup /topic groupname
tatcli topic deleteAdminGroup /topic groupname1 groupname2
```

### Communicate with Tat Websocket

#### Interactive mode
```
tatcli socket interactive
```

#### Subscribe to messages event on one topic

Returns a full message struct:

```
tatcli socket messages <<topic>|all>> [topic]...
```

Returns a full message struct, with replies:

```
tatcli socket messages <<topic>|all|onetree|fulltree>> [topic]...
```

Example:

```
tatcli socket messages onetree all
tatcli socket messages fulltree all
tatcli socket messages fulltree /TopicA/subTopicA
```

#### Subscribe to new messages event on one or many topic
Returns only topic name when a new message is created on this topic
```
tatcli socket messagesNew <<topic>|all>> [topic]...
```

#### Write a presence on one topic
```
tatcli socket writePresence <status> <topic>...
```

#### Subscribe to presences event on one topic
```
tatcli socket presences <<topic>|all>> [topic]...
```

#### Subscribe to user event, for admin only
```
tatcli socket users
```


# Hacking

Tatcli is written in Go 1.5, using the experimental vendoring
mechanism introduced in this version. Make sure you are using at least
version 1.5.

```bash
mkdir -p $GOPATH/src/github.com/ovh
cd $GOPATH/src/github.com/ovh
git clone git@github.com:ovh/tatcli.git
cd $GOPATH/src/github.com/ovh/tatcli
export GO15VENDOREXPERIMENT=1
go build
```

You've developed a new cool feature? Fixed an annoying bug? We'd be happy
to hear from you! Make sure to read [CONTRIBUTING.md](./CONTRIBUTING.md) before.
