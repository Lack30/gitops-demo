FROM scratch

COPY . .

EXPOSE 8080

CMD ["nohup ./main &"]
