# What purpose serves the python script ?

This python script in used to populate the database with the minimum essential data coming from the base .csv file

# What structure needs the .csv file to file ?

The base .csv is composed of 15 columns in the following order :

- 1. Event : The type of event the match was played on
- 2. White : White's player's username
- 3. Black : Black's player's username
- 4. Result : Outcome of the game
- 5. UTCDate : Date when the game happened
- 6. UTCTime : Time when the game happened
- 7. WhiteElo : The elo rating of white's player
- 8. BlackElo : The elo rating of black's player
- 9. WhiteRatingDiff : The elo gained or lost after the game
- 10. BlackRatingDiff : The elo gained or lost after the game
- 11. ECO : ECO encoding of the name of the opnening
- 12. Opening : Opening name
- 13. TimeControl : Time of the game in secondes with additional time control
- 14. Termination : How did the game end (Normal, Forfeit, ...)
- 15. AN : Movements in Movetext format

The base csv does not include headers. You need to use indexes to go through the csv.