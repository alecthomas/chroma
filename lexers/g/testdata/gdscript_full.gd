class_name StateMachine
extends Node

signal state_changed(previous, new)

export var initial_state = NodePath()

func _input(event):
    if event.is_action_pressed("jump"):
        jump()
