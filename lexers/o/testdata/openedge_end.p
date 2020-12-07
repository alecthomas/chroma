ON CHOOSE OF send-button DO:
  RUN runRemoteProc.
  S1 = "Ran proc(" + STRING(xmtcnt) + ")".
  DISPLAY S1 WITH FRAME foo 1 DOWN.
  HIDE FRAME bar.
END.
