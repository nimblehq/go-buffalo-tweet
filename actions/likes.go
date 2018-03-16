package actions

import (
    "github.com/bufftwitt/models"
    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/pop"
    "github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Like)
// DB Table: Plural (likes)
// Resource: Plural (Likes)
// Path: Plural (/likes)
// View Template Folder: Plural (/templates/likes/)

// LikesResource is the resource for the Like model
type LikesResource struct {
    buffalo.Resource
}

// List gets all Likes. This function is mapped to the path
// GET /likes
func (v LikesResource) List(c buffalo.Context) error {
    // Get the DB connection from the context
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return errors.WithStack(errors.New("no transaction found"))
    }

    likes := &models.Likes{}

    // Paginate results. Params "page" and "per_page" control pagination.
    // Default values are "page=1" and "per_page=20".
    q := tx.PaginateFromParams(c.Params())

    // Retrieve all Likes from the DB
    if err := q.All(likes); err != nil {
        return errors.WithStack(err)
    }

    // Make Likes available inside the html template
    c.Set("likes", likes)

    // Add the paginator to the context so it can be used in the template.
    c.Set("pagination", q.Paginator)

    return c.Render(200, r.HTML("likes/index.html"))
}

// Show gets the data for one Like. This function is mapped to
// the path GET /likes/{like_id}
func (v LikesResource) Show(c buffalo.Context) error {
    // Get the DB connection from the context
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return errors.WithStack(errors.New("no transaction found"))
    }

    // Allocate an empty Like
    like := &models.Like{}

    // To find the Like the parameter like_id is used.
    if err := tx.Find(like, c.Param("like_id")); err != nil {
        return c.Error(404, err)
    }

    // Make like available inside the html template
    c.Set("like", like)

    return c.Render(200, r.HTML("likes/show.html"))
}

// Create adds a Like to the DB. This function is mapped to the
// path POST /likes
func (v LikesResource) Create(c buffalo.Context) error {
    // Allocate an empty Like
    like := &models.Like{}

    // Bind like to the html form elements
    if err := c.Bind(like); err != nil {
        return errors.WithStack(err)
    }

    // Get the DB connection from the context
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return errors.WithStack(errors.New("no transaction found"))
    }

    // Validate the data from the html form
    verrs, err := tx.ValidateAndCreate(like)
    if err != nil {
        return errors.WithStack(err)
    }

    if verrs.HasAny() {
        // Make like available inside the html template
        c.Set("like", like)

        // Make the errors available inside the html template
        c.Set("errors", verrs)

        // Render again the new.html template that the user can
        // correct the input.
        return c.Render(422, r.HTML("likes/new.html"))
    }

    // If there are no errors set a success message
    //c.Flash().Add("success", "Like was created successfully")

    // and redirect to the same All Tweets page
    return c.Redirect(302, "/all_tweets")
}

// Edit renders a edit form for a Like. This function is
// mapped to the path GET /likes/{like_id}/edit
func (v LikesResource) Edit(c buffalo.Context) error {
    // Get the DB connection from the context
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return errors.WithStack(errors.New("no transaction found"))
    }

    // Allocate an empty Like
    like := &models.Like{}

    if err := tx.Find(like, c.Param("like_id")); err != nil {
        return c.Error(404, err)
    }

    // Make like available inside the html template
    c.Set("like", like)
    return c.Render(200, r.HTML("likes/edit.html"))
}

// Update changes a Like in the DB. This function is mapped to
// the path PUT /likes/{like_id}
func (v LikesResource) Update(c buffalo.Context) error {
    // Get the DB connection from the context
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return errors.WithStack(errors.New("no transaction found"))
    }

    // Allocate an empty Like
    like := &models.Like{}

    if err := tx.Find(like, c.Param("like_id")); err != nil {
        return c.Error(404, err)
    }

    // Bind Like to the html form elements
    if err := c.Bind(like); err != nil {
        return errors.WithStack(err)
    }

    verrs, err := tx.ValidateAndUpdate(like)
    if err != nil {
        return errors.WithStack(err)
    }

    if verrs.HasAny() {
        // Make like available inside the html template
        c.Set("like", like)

        // Make the errors available inside the html template
        c.Set("errors", verrs)

        // Render again the edit.html template that the user can
        // correct the input.
        return c.Render(422, r.HTML("likes/edit.html"))
    }

    // If there are no errors set a success message
    c.Flash().Add("success", "Like was updated successfully")

    // and redirect to the likes index page
    return c.Redirect(302, "/likes/%s", like.ID)
}

// Destroy deletes a Like from the DB. This function is mapped
// to the path DELETE /likes/{like_id}
func (v LikesResource) Destroy(c buffalo.Context) error {
    // Get the DB connection from the context
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return errors.WithStack(errors.New("no transaction found"))
    }

    // Allocate an empty Like
    like := &models.Like{}

    // To find the Like the parameter like_id is used.
    if err := tx.Find(like, c.Param("like_id")); err != nil {
        return c.Error(404, err)
    }

    if err := tx.Destroy(like); err != nil {
        return errors.WithStack(err)
    }

    // If there are no errors set a flash message
    c.Flash().Add("success", "Like was destroyed successfully")

    // Redirect to the likes index page
    return c.Redirect(302, "/likes")
}
