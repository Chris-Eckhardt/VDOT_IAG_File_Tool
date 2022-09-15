# VDOT_IAG_File_Tool

This is a command line tool I made to help with searching VDOT ITAG or ITGU files.

It takes 3 parameters: -path, -param, -type

## avilookup -type id -path C:\wherever\0010_0923432432.ITAG -param 0000000037

It is searching for tags with an id matching 0000000037.

It will return all tags found in the -path file that match the -type and -param criteria. Be careful to not be too broad because it can return A LOT of results. Best to stick with specifics like ID.

Parameter order does not matter.

There are 3 -type arguments that are allowed: 
•	id (10 characters),
•	class (IAG class and 4 characters, eg. 0072 or ****), 
•	protocol (3 characters, eg. T  , ***)

-path can be relative or absolute path to the ITAG or ITGU file.

-param is the argument corresponding to the -type command.

All arguments for param should meet IAG spec in order for the parser to find them.

If you plan on using this tool a lot, you can add it to your path in system variables so it can be called from anywhere.
