library(readr)
library(dplyr)
library(ggplot2)
library(maps)

cityfile <- "gb_cities.csv"
routefile <- "route.txt"

cities <- read_csv(cityfile, show_col_types = FALSE)
route <- read_csv(routefile, show_col_types = FALSE) %>%
  pull(route)
n <- length(route) # nr cities
route <- route + 1 # for 1-up indexing in R

# order dataframe by route
cities <- cities[route,]

toLat <- c(cities$Latitude[2:n], cities$Latitude[1])
toLong <- c(cities$Longitude[2:n], cities$Longitude[1])
cities['toLat'] <- toLat
cities['toLong'] <- toLong

# draw cities
worldmap <- map_data('world')
ggplot() + 
  geom_polygon(data = worldmap, 
               aes(x = long, y = lat, 
                   group = group), 
               fill = 'gray90', 
               color = 'black') + 
  coord_fixed(ratio = 1.3, 
              xlim = c(-10,3),   # lat/long limits for UK
              ylim = c(50, 59)) + 
  theme_void() +
  # draw route
  geom_segment(data = cities,
               aes(x = Longitude, y = Latitude, 
                   xend = toLong, yend = toLat), 
               col = 'blue') +
  # draw cities
  geom_point(data = cities, 
               aes(x = Longitude, y = Latitude),
             col='red') +
  geom_text(data = cities,
            aes(x = Longitude, y = Latitude, label = `Place Name`),
            size = 2,
            nudge_x = 0.3,
            check_overlap = TRUE)

  

