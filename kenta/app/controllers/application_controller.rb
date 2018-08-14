class ApplicationController < ActionController::API
  include Knock::Authenticable

  private
  def admin_user?
    if current_user.status.admin?
      render json: { text: "reject" }
    end
  end
end
