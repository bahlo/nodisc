package controllers

import "github.com/revel/revel"

func init() {
  revel.OnAppStart(InitDB)
  revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
  revel.InterceptMethod(App.AddUser, revel.BEFORE)
  revel.InterceptMethod(App.AddConfig, revel.BEFORE)
  revel.InterceptMethod(Threads.checkUser, revel.BEFORE)
  revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
  revel.InterceptMethod((*GorpController).Rollback, revel.PANIC)

  revel.TemplateFuncs["eqo"] = func(a, b, c interface{}) bool {
    return a == b || a == c
  }
}
