This is the documentation.

## Getting Started

My weapon of choice is IdeaJ with the Golang extensions. It is not great but sufficient, the best you can get if you used to an IDE.
Vim is an editor :)

## TODO List

- [ ] Write Todo
- [ ] Remove Write Todo
- [ ] Pick a config format, write a parser

## Terms

### Types
* ***Funnel*** a single regexp that on match notifies an a ***Action***
* ***FunnelGroup*** a collection of funnels watching the same stream

### Instances
* ***Tracker*** an instance of FunnelGroup tracking a specific stream marked by a ***SessionId***
* ***Action*** if a ***Tracker***'s Funnel got a match it notifies the ***Action***