# call with 
# Rscript drawTree.R newick_file

library(phytools)

args <- commandArgs(trailingOnly=T)

if(length(args) < 1){
  cat("Usage: drawTree.R newick_file\n")
  quit(status=1)
}
newickFile <- args[1]

# e.g.
# newickFile <- "newick.txt"

# read file
tree <- read.tree(newickFile)
count <- tree$edge.length

# adjust parameters
eLength <- 2
tree$edge.length <- rep(eLength, nrow(tree$edge))

nedges <- nrow(tree$edge)
ntips <- length(tree$tip.label)
nnodes <- length(tree$node.label)
ncount <- rep(1, nnodes)
tcount <- rep(1, ntips)

for(i in 1:nedges){
  k <- tree$edge[i,2] # find endpoint
  if(k <= ntips){
    # endpoint is a tip (i.e. a leaf)
    tcount[k] <- count[i]
  } else {
    # endpoint is in internal node
    ncount[k - ntips] <- count[i]
  }
}

ctmax <- max(c(tcount, ncount))
tcex <- (4*tcount + ctmax - 5)/(ctmax - 1)
ncex <- (4*ncount + ctmax - 5)/(ctmax - 1)

# plot tree
plot(tree,
     type='cladogram',
     no.margin=TRUE,
     edge.width=1,
     edge.color="grey",
     tip.color="blue",
     cex=0.5*tcex)
nodelabels(text=tree$node.label,
           node=1:tree$Nnode+Ntip(tree),
           frame="none",
           adj=c(1.1,-0.4),
           col="blue",
           cex=0.5*ncex)

