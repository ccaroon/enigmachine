from invoke import Collection, task


@task
def test(ctx, verbose=False):
    """Run Unit Tests"""
    flags = []
    if verbose:
        flags.append("-v")

    # With Ginkgo
    ctx.run(f"ginkgo {' '.join(flags)} ./... ")

@task(
    aliases=["cover"]
)
def coverage(ctx):
    """Run Code Coverage"""

    # With Ginkgo
    ctx.run("ginkgo -cover ./...")
    ctx.run("go tool cover -html coverprofile.out -o coverage.html")

ns = Collection(
    test,
    coverage
)
