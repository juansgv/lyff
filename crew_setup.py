from crewai import Agent, Crew, Task, Process
from crewai_tools import FileReadTool
from config import OPENAI_API_KEY

# Set the OpenAI API key
import os
os.environ['OPENAI_API_KEY'] = OPENAI_API_KEY

# Define the Data Digest Agent
data_digest_agent = Agent(
    role='Data Digest Agent',
    goal='Digest and summarize company data from Excel files',
    backstory="An AI agent responsible for securely reading and summarizing company Excel data.",
    tools=[FileReadTool()],
    verbose=True
)

# Define the Analysis Agent
analysis_agent = Agent(
    role='Analysis Agent',
    goal='Analyze the summarized data to provide actionable recommendations',
    backstory="An AI agent that analyzes summarized data to offer strategic improvement suggestions.",
    tools=[],
    verbose=True
)

# Define tasks
digest_task = Task(
    description='Read and summarize the contents of the uploaded Excel file.',
    expected_output='A summary of the company\'s state based on the Excel data.',
    agent=data_digest_agent
)

analysis_task = Task(
    description='Analyze the summarized data to generate actionable recommendations for company improvement.',
    expected_output='A comprehensive report with strategic recommendations.',
    agent=analysis_agent
)

# Assemble the crew
crew = Crew(
    agents=[data_digest_agent, analysis_agent],
    tasks=[digest_task, analysis_task],
    process=Process.sequential,
    memory=True,
    verbose=True
)