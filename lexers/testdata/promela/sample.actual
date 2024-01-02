/**
	Needham-Schroeder message passing protocol (patched). 
	|	msg1:	agentA -> agentB	(keyB, agentA, nonceA, 0)
	|	msg2:	agentB -> agentA	(keyA, agentB, nonceA, nonceB)
	|	msg3:	agentA -> agentB	(keyB, nonceB, 0, 0)

	Note that sending (keyB, agentA, nonceA) from agentA to agentB 
	over the network (chan)nel models agentA encrypting the message 
	"[agentA, nonceB]" with agentB's public key.
*/

mtype = {
	/* Status Codes */
	ok, 
	err, 

	/* Message Codes */
	msg1, 
	msg2, 
	msg3, 

	/*	Agent (A)lice */
	keyA,
	agentA, 
	nonceA,

	/*	Agent (B)ob */
	keyB, 
	agentB,
	nonceB,

	/*	Agent (I)ntruder */
	keyI, 
	agentI, 
	nonceI 
};


/**
	Declare a structured data type to model our encrypted messages.
	Messages will contain either 2 or 3 content fields.
*/
typedef Crypt { 
	mtype key, 
	content1, 
	content2,
	content3 
};


/**
	Model network between agents via a rendezvous channel. 
	Send and recieve operations are performed synchronously. 
*/
chan network = [0] of {mtype, mtype, Crypt};


/* global variables for verification*/
mtype partnerA; 
mtype partnerB;
mtype statusA = err;
mtype statusB = err;
bool knows_nonceA; 
bool knows_nonceB;

/* Agent (A)lice */
active proctype Alice() {

	/* local variables */
	mtype pkey;			/* reciever's public key */
	mtype pnonce;		/* reciever's nonce	 */
	Crypt messageAB;	/* sent messages			 */
	Crypt data;			/* recieved messages	 */

	/* 
		Initialization: In this example we non-deterministically choose between 
		agents (B)ob and (I)ntruder
	*/
	if 
	:: partnerA = agentB; pkey = keyB;
	:: partnerA = agentI; pkey = keyI;
	fi;

	/* prepare (msg1) */
	messageAB.key = pkey;
	messageAB.content1 = agentA;
	messageAB.content2 = nonceA;
	messageAB.content3 = 0;

	/* send (msg1) */
	network ! msg1 (partnerA, messageAB);

	/* recv (msg2) : blocking */
	network ? (msg2, agentA, data);

	/* verify (msg2) : blocking	*/
	(data.key == keyA) && (data.content1 == partnerA) && (data.content2 == nonceA);

	/* obtain (msg2) sender's nonce */
	pnonce = data.content3;

	/* prepare (msg3) */
	messageAB.key = pkey;
	messageAB.content1 = pnonce;
	messageAB.content2 = 0;
	messageAB.content3 = 0;

	/* send (msg3) */
	network ! msg3 (partnerA, messageAB);

	/* update status */
	statusA = ok;
}

/* Agent (B)ob */
active proctype Bob() {
		
	/* local variables */
	mtype pkey;			/* reciever's public key */
	mtype pnonce;		/* reciever's nonce	 */
	Crypt messageBA;	/* sent messages			 */
	Crypt data;			/* recieved messages	 */

	/* Initialization	*/
	partnerB = agentA;
	pkey	 = keyA;

	/* recv (msg1) : blocking */
	network ? (msg1, agentB, data)

	/* verify (msg1) : blocking */
	(data.key == keyB) && (data.content1 == agentA);

	/* obtain (msg1) sender's nonce */
	pnonce = data.content2;	

	/* prepare (msg2) */
	messageBA.key = pkey;
	messageBA.content1 = agentB;
	messageBA.content2 = pnonce;
	messageBA.content3 = nonceB;

	/* send (msg2) */
	network ! msg2 (partnerB, messageBA);

	/* recv (msg3) : blocking */
	network ? (msg3, agentB, data);

	/* verify (msg3) : blocking */
	(data.key == keyB) && (data.content1 == nonceB);

	statusB = ok;
}

/* Agent (I)ntruder */
active proctype Intruder() {

	mtype msg, recpt;
	Crypt data, intercepted;

	/* Initialize knows_nonce variables to false */
	knows_nonceA = false;
	knows_nonceB = false;

	do
	:: network ? (msg, _, data) -> 
	if /* perhaps store the message */
			
		::	intercepted.key		 = data.key;
			intercepted.content1 = data.content1;
			intercepted.content2 = data.content2;
			
			/*	
				Message contains (I)ntruder's (public) key, the intruder can 
				decrypt the message. Note that we can learn nonce values from 
				either content1 or content2
				|	msg1:	(keyB, agentA, nonceA, 0)
				|	msg2:	(keyA, agentB, nonceA, nonceB)
				|	msg3: 	(keyB, nonceB, 0, 0)
			*/
			if 
			::  (intercepted.key == keyI) -> if 
				:: intercepted.content1 == nonceA -> knows_nonceA = true; 
				:: intercepted.content1 == nonceB -> knows_nonceB = true; 
				:: intercepted.content2 == nonceA -> knows_nonceA = true; 
				:: intercepted.content2 == nonceB -> knows_nonceB = true; 
				:: intercepted.content3 == nonceA -> knows_nonceA = true; 
				:: intercepted.content3 == nonceB -> knows_nonceB = true;
				fi 

			:: skip;
			fi 

		:: skip;
	fi;

	:: /* Replay or send a message */
	if /* choose message type */
		:: msg = msg1;
		:: msg = msg2;
		:: msg = msg3;
	fi;
		
	if /* choose a recepient */
		:: recpt = agentA;
		:: recpt = agentB;
	fi;
			 
	if /* replay intercepted message or assemble it */
		::	data.key		= intercepted.key;
			data.content1	= intercepted.content1;
			data.content2	= intercepted.content2;
			data.content3	= intercepted.content3;
		
		:: 
		if /* assemble content1 */
			:: data.content1 = agentA;
			:: data.content1 = agentB;
			:: data.content1 = agentI;
			:: (knows_nonceA) -> data.content1 = nonceA
			:: (knows_nonceB) -> data.content1 = nonceB
			:: (!knows_nonceA && !knows_nonceB) -> data.content1 = nonceI;
		fi;
		
		if /* assemble key */
			:: data.key = keyA;
			:: data.key = keyB;
			:: data.key = keyI;
		fi;
		
		if 
		:: (knows_nonceA) -> data.content2 = nonceA
		:: (knows_nonceB) -> data.content2 = nonceB
		:: (!knows_nonceA && !knows_nonceB) -> data.content2 = nonceI;
		fi 

		if 
		:: (knows_nonceA) -> data.content3 = nonceA
		:: (knows_nonceB) -> data.content3 = nonceB
		:: (!knows_nonceA && !knows_nonceB) -> data.content3 = nonceI;
		fi 
		
	fi;

	network ! msg (recpt, data);
	od
}

/**
	Always, one process will terminate in error
*/
ltl alwaysErr { [] ( (statusA == err) || (statusB == err) ) }

/**
	Eventually the protocol will complete without error
*/
ltl eventuallyOk { <> ( (statusA == ok) && (statusB == ok) ) }

/*
	propAB: If both Alice and Bob reach the end of their runs (i.e. both statusA and statusB are ok) 
	then Alice's communication partner is Bob, and Bob's communication partner is Alice.
*/
ltl propAB { [] ( ( (statusA == ok) && (statusB == ok) ) -> ( (partnerB == agentA) && (partnerA == agentB) ) ) }

/*
	propA: If agent A reaches the end of its run (statusA is ok) and A believes it is talking to B 
	(partnerA is agentB) then the intruder does not know A's nonce (!knows_nonceA).
*/
ltl propA { [] ( ( (statusA == ok)  && (partnerA == agentB) ) ->  ( knows_nonceA == false ) ) }

/*
	propB: If agent B reaches the end of its run (statusB is ok) and it believes it is talking to A 
	(partnerB is agentA) then the intruder does not know B's nonce (!knows_nonceB)
*/
ltl propB { [] ( ( (statusB == ok)  && (partnerB == agentA) ) ->  ( knows_nonceB == false ) ) }
