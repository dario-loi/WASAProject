# WasaPhoto API

## Introduction

We want to spend a few words on what the rationale behind WASAPhoto will be, since we 
think that from this moment on either we plan this out perfectly or we fail spectacularly.

we identify the criticalities:

- We don't handle photos uniformly at ALL
- The resources poorly represent atomic entities (users)

All of this prevents the concretization of a good API, we **need** to fix it

## The goals

We want to provide

1. a uniform way to handle photos
2. a minimal rethinking of the API in order to directly extract the necessary resources for frontend components.

## Photo Handling

Photos are present in our application in a number of ways:

- As a profile picture
  - In the profile page
  - Embedded in a little circle near the username
- As part of a Stream
  - In the streams page
  - In a user's gallery

Across all of these interfaces, photo loading is almost always done as a part of an 
`application/json` request, we want to split this into two stages:

The first stage is the obtaining of the photo's UUID and all related metadata, this allows 
us to separate the heavy load of the image from the rest

The second stage is to load the photo using the UUID, this is a shared endpoint that will 
be requested at a separate time.

#### The takeaway

This means that the *only* endpoint to return a photo will be the one that accepts 
an ID and retrieves the photo, all other endpoints will return the UUID and metadata 
which will be used to access this single interface.

## Resources & The Frontend

Now, we imagine WASAPhoto to be not too dissimilar to the social network Instagram,
this in turn means that we imagine the frontend to be composed of a number of components

1. The profile page, containing:
   1. A Gallery
   2. A Biography
   3. Followers/Following Posts counters
2. A Navbar to search for other users
3. A feed (we refer to it as **Stream**)
   - Each feed is composed of PhotoPosts
     - Each photopost is composed of:
       1. The photo itself
       2. Author name
       3. Like counter
       4. Comments, made up of
          1. The commenter's profile photo/name
          2. comment body 


The most complex entity in the bunch is the *Photopost*, the whole task seems to be 
**daunting** but we quickly need a way to tackle this lest we be overwhelmed.

## The plan

We need an easy step by step plan to follow:

- [ ] Create a new endpoint to retrieve a photo by UUID
- [ ] Standardize all other photo occurrences to return UUID and metadata and al other components to behave as such
- [ ] Store up photos in the filesystem in the backend in a directory such as `./photos/{UUID}.jpg`

We will *then* proceed to tackle the frontend, which will be a separate issue.
- [ ] Standardize the other endpoints to return more *useful* resources.



