# API Routes

## Events

- **GET** `/events` → Get a list of all available events
- **GET** `/events/[id]` → Get event with this ID
- **POST** `/events` → Create a new bookable event _(auth required)_
- **PUT** `/events/[id]` → Update an event _(auth required, only by creator)_
- **DELETE** `/events/[id]` → Delete an event _(auth required, only by creator)_

## Authentication

- **POST** `/signup` → Create a new user
- **POST** `/login` → Login as a user _(returns auth token)_

## Event Registration

- **POST** `/events/[id]/register` → Register a user for an event _(auth required)_
- **DELETE** `/events/[id]/register` → Cancel registration _(auth required)_
