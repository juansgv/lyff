import os
from dotenv import load_dotenv

# Load environment variables from .env file
load_dotenv()

# API Keys
OPENAI_API_KEY = os.getenv('OPENAI_API_KEY')
# Add any other API keys you might need, for example:
# SERPAPI_API_KEY = os.getenv('SERPAPI_API_KEY')

# You can add other configuration variables here as needed