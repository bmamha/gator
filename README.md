#Project: gator 
#A blog aggregator project for boot.dev

### Description
gator is a blog aggregator CLI program built in go and using PostgreSQL.

### Requirements 

go version 1.2.3 
PostgreSQL 16.6 

### How to Install gator:

1. Clone project 
2. Manually Create .gatorconfig.json file into your repository 
And add the following JSON file
'''
{
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}
'''
3. Create a new database called `gator` using PostgreSQL

4. Run `go install` on your terminal 


### How to use program 

The following commands can be run using gator.

1. `gator register <name>` To register a user for the program
2. `gator login <name>` To login as a user for the program.
3. `gator reset` To delete all data stored in the database.
4. `gator users` To display all users in the database
5. `gator addfeed <url>` To add a website feed for the user who is logged in.
6. `gator feeds` Displays the blog feeds added to our program.
7. `gator follow <url>` Allows user to follow a feed if it has already been added to database.
8. `gator following` Displays the blogs followed by current(logged in) user.
9. `gator unfollow <url>` Allows user to unfollow a feed.
10. `gator agg <time(e.g. 1s)` Fetches  and adds to database posts from the blogs followed by current user. The time parameter controls how we frequently websites are fetched.  
11. `gator browse ` Displays posts from the feeds followed by user.

