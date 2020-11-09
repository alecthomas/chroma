Theorem demorgan : forall (P Q : Prop),
  ~(P \/ Q) -> ~P /\ ~Q.
Proof.
  tauto.
Qed.
