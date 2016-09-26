FROM scratch
ADD redirect /
ENTRYPOINT ["./redirect"]