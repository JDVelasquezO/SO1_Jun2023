# Desde que lenguaje o framework queremos trabajar
FROM node:18

# Crear nuestro directorio de trabajo
WORKDIR /api

# Copiar los archivos necesarios para ejecutar el código
COPY package.json ./
COPY package-lock.json ./

# Ejecutar el comando para obtener librerías
RUN npm install

# Llamar variables de entorno
ENV DB_HOST localhost

# Copiar el resto del código que se va a ejecutar
COPY . .

# Ejecutar el comando para lanzar la aplicación
CMD ["npm", "start"]