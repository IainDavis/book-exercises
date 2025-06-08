from is_unique import is_unique_for_with_set, is_unique_leader_follower, is_unique_bit_vector
import pytest 

implementations = [
    is_unique_for_with_set,
    is_unique_leader_follower,
    is_unique_bit_vector
]

@pytest.mark.parametrize("impl", implementations)
def test_unique_simple(impl):
    assert impl("abc") is True
    assert impl("aabc") is False

@pytest.mark.parametrize("impl", implementations)
def test_unique_unicode(impl):
    assert impl("🙉💩") is True
    assert impl("🙉🙉") is False

@pytest.mark.parametrize("impl", implementations)
def test_empty_input(impl):
    assert impl("") is True

@pytest.mark.parametrize("impl", implementations)
def test_none_input(impl):
    with pytest.raises(TypeError):
        assert impl(None)
        
@pytest.mark.parametrize("impl", implementations)
def test_non_string_input(impl):
    with pytest.raises(TypeError):
        assert impl(1.1)
        assert impl(1)
        assert impl(1, 2, 3)
        assert impl([1, 2, 3])
        assert impl({1, 2, 3})
        assert impl({1: "value", 2: "value", 3: "value"})

