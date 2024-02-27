# Go Auto Move Files
## Automation for organizing files and folders in the user OS.

### The Problem
- People who organize several files in their systems might spend too much time saving the same docs and monthly bills from their Downloads folder to the desired folder. This project aims the save time by reading some specific key in the file name and moving it automatically to the user desired path.
- Once the user set the desired paths in the json, it will only have the burden of when saving a file, in the moment he will name it, he will have to add the key name.

- It is possible to analyze the content and headers of the file, but that will be for the future.

- Warning: Paths must match the path names shown in the terminal, translated paths shown in the gui might not work. 
- Warning: For windows users, don't use backslashes in json, use forward slashes "/".

### Options:
- 1 **Default**: Simply move file x to destination y.
- 2 **With Date**: Get file date and create/move to folder accordingly to the year/month/day *Not yet implemented*
  - Configs:
    - Daily: will save in the following folder: key/year/month/day
    - Monthly: will save in the following folder: key/year/month
    - Yearly: will save in the following folder: key/year
- 3 **Delete**: Simply delete file from src path. Useful for some cases where extensions automatically write in the Downloads folder since they can't have access to user OS. *Not yet implemented*

roadmap:
*with date*
*delete*
*support to folders not yet implemented*
*add retry operations*
*tests*
