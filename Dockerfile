FROM busybox

COPY ./bin/serviced /serviced

CMD ["/serviced"]
