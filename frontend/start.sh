#!/bin/sh
echo "REACT_APP_IMAGE_VIEWER_BACKEND_HOST=${IMAGE_VIEWER_BACKEND_HOST}" > .env

# Start server
npm run start