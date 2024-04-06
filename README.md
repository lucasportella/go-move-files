# Go Auto Move Files
## Automation for organizing files and folders in the user OS.

This is an on going project,
roadmap:
*add interfaces*
*tests*
*support to folders*
*add retry operations*
*add automated service to run project example*

### The Problem
- People who organize several files in their systems might spend too much time saving the same docs and monthly bills from their Downloads folder to the desired folder. This project aims to save time by reading some specific key in the file name and moving it automatically to the user desired path.
- Once the user set the desired paths in the json, it will only have the burden of when saving a file, in the moment he will name it, he will have to add the key name.

- It is possible to analyze the content and headers of the file, but that will be for the future.

- Warning: Paths must match the path names shown in the terminal, translated paths shown in the gui might not work. 
- Warning: For windows users, don't use backslashes in json, use forward slashes "/".


### How to use
- 1 create a JSON file named *paths.json* and follow the example of *paths-example.json*
- 2 change the keys #book, #music, #picture, #bill, #doc etc to the keys that will be in your filename, In the *paths-example.json* any filename with #book will go to the book path, any filename with #bill will got to the bill path + /current_year/current_month

### Options:
- 1 **Default**: Simply move file x to destination y.
- 2 **With Date**: Get file date and create/move to folder accordingly to the year/month/day 
  - Configs:
    - Daily: will save in the following folder: key/year/month/day
    - Monthly: will save in the following folder: key/year/month
    - Yearly: will save in the following folder: key/year


