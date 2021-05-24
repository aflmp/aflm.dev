<article>
    <p>Why am I creating a blog which houses static pages from scratch? Isn't this trivially done if one were to use static site generators?
        To answer that - Yes, it's kind of like reinventing the wheel but I'm doing this to with multiple things in mind - to explore net/http package, to explore cloud run for personal use, 
        and to finally have a blog that is all written by me. I wanted something simple that does not require a lot of dependencies and maintainence. The outcome is a simple blog built out of go standard library alone.
        <p>This turned out to be a good learning project and would continue to evolve as needed.
        If you are planning to create a blog for yourself, I'd suggest trying out jekyll, hugo and others first.
        This is a learning projectThis fIn this post let's talk about how this blog is built and deployed. This blog is a learning project, If you are looking to creaBefore we begin, If you are looking to create blog
        for yourself, you can> </p>
    </p>
    <p>It is built using go, with no external dependencies. The source for the blog lives 
        <a href="https://github.com/aflmp/" target="_blank"></a>here.

        

    </p> 
    <h4>Project Structure</h4>
    <p>
        <pre class="code">
            ~/aflm.dev
            ❯ tree -L 1 -d
            .
            ├── assets
            ├── build
            ├── cmd
            ├── pages
            ├── posts
            └── templates
            
            6 directories
        </pre>
        The project follows the <a href="" target="_blank">standard go project layout</a>.</p>

        <p><code>posts.json</code> is a file used for recording each post on the blog. I initially used a yaml file which 
        added to the only external dependency for this project as the standard library lacks one. I don't like json for
        it's verbosity but it was the next best option.
    </p>

    How do you write posts?
    - write them in HTML
    - create an entry for the same in posts.config
    - push to the master branch

    How do you deploy?
    - the blog is deployed as part of github actions on the repository
    - it involves building a docker image from the dockerfile
    - pushing the dockerimage to the gcr registry 
    - deploying the cloud run application using service account



</article>