.method public getTokens(I)I
    .locals 2
    .param p1, "amt"    # I

    .prologue
    const/4 v0, 0x0

    .line 2
    iget-boolean v1, p0, Lcom/limbenjamin/Example;->isPaid:Z

    if-nez v1, :cond_1

    .line 5
    :cond_0
    :goto_0
    return v0

    .line 2
    :cond_1
    iget-object v1, p0, Lcom/limbenjamin/Example;->handler:Lcom/limbenjamin/ExampleHandler;

    if-eqz v1, :cond_0

    .line 3
    move v3, p1

    iget-object v0, p0, Lcom/limbenjamin/Example;->handler:Lcom/limbenjamin/ExampleHandler;

    invoke-interface {v0, v3}, Lcom/limbenjamin/ExampleHandler;->creditTokens(I)V

    move-result v0

    goto :goto_0
.end method
