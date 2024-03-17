monthConversions = {
    "Jan":"January",
    "Feb":"February",
    "Mar":"March",
    #key value pairs
    #key must be unique

}




print(monthConversions["Feb"])
print(monthConversions.get("Jan"))


print(monthConversions.get("Luv","not a valid"))