# Используем базовый образ Nginx
FROM nginx:alpine

# Копируем кастомный конфигурационный файл Nginx
COPY nginx.conf /etc/nginx/nginx.conf

# Копируем файлы из папки Pages в директорию Nginx
COPY pages/ /usr/share/nginx/html/pages/

# Открываем порт 80 (стандартный для Nginx)
EXPOSE 80

# Запускаем Nginx
CMD ["nginx", "-g", "daemon off;"]
