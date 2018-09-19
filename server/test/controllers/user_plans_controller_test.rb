require 'test_helper'

class UserPlansControllerTest < ActionDispatch::IntegrationTest
  test "should get index" do
    get user_plans_index_url
    assert_response :success
  end

end
