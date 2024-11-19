FROM nginx:alpine

COPY nginx.conf /etc/nginx/nginx.conf
COPY pages/ /usr/share/nginx/html/pages/
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
