module VerseRoleTest
  def test_plays_verse_role
    assert_respond_to @role_player, :lyrics
  end
end
