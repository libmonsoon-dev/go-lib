package com.github.libmonsoondev.golib.ctxbind

import ctxbind.Ctxbind
import org.junit.Assert
import org.junit.Test

class CtxBindTest {
    @Test
    fun throwsAfterClosed() {
        val ctxID = Ctxbind.withCancel(Ctxbind.backgroundContext())

        Ctxbind.err(ctxID)
        Ctxbind.cancelContext(ctxID)

        val gotErr = Assert.assertThrows(Ctxbind.getContextCanceled().javaClass) { Ctxbind.err(ctxID) }
        Assert.assertTrue(Ctxbind.isContextCanceled(gotErr))
    }

    //TODO: add more tests
}
