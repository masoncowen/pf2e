from run_session import Context, Command, States

def test_allowed_states():
  cont = Context(state = States.Exploration)
  cont2 = Context(state = States.Combat)
  def func():
    pass
  comm = Command(text = '', description = '', function = func,
                 allowed_states={States.Exploration})
  comm2 = Command(text = '', description = '', function = func,
                 allowed_states={States.Combat})
  assert comm.can_run_in_context(cont)
  assert not comm2.can_run_in_context(cont)
  assert comm2.can_run_in_context(cont2)
  assert not comm.can_run_in_context(cont2)
