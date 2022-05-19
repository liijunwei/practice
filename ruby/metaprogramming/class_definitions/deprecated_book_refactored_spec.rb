require 'rspec'
require_relative 'deprecated'

RSpec.describe BookRefactored do
  describe "#GetTitle" do
    it "responds to GetTitle" do
      expect(subject.GetTitle).to eq("Different Seasons")
    end

    it "responds to title" do
      expect(subject.title).to eq("Different Seasons")
    end
  end

  describe "#title2" do
    it "responds to title2" do
      expect(subject.title2).to eq("Rita Heyworth and Shawshank Redemption")
    end

    it "responds to subtitle" do
      expect(subject.subtitle).to eq("Rita Heyworth and Shawshank Redemption")
    end
  end

  describe "#LEND_TO_USER" do
    it "responds to LEND_TO_USER" do
      expect(subject.LEND_TO_USER("Bob")).to eq("Lending Different Seasons to Bob")
    end

    it "responds to lend_to" do
      expect(subject.lend_to("Bob")).to eq("Lending Different Seasons to Bob")
    end
  end
end
