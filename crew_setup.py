import os
import pandas as pd
from crewai import Agent, Crew, Task, Process
from crewai.tools import CrewAITool
from langchain.tools import Tool
from config import OPENAI_API_KEY

# Set the OpenAI API key
os.environ['OPENAI_API_KEY'] = OPENAI_API_KEY

def read_tabular_data(file_path):
    """
    Reads tabular data from various file formats including Excel and CSV.

    Args:
        file_path (str): The path to the data file.

    Returns:
        pd.DataFrame: The loaded data as a pandas DataFrame.
    """
    try:
        if file_path.endswith('.xlsx') or file_path.endswith('.xls'):
            return pd.read_excel(file_path)
        elif file_path.endswith('.csv'):
            return pd.read_csv(file_path)
        else:
            raise ValueError(f"Unsupported file format: {file_path}")
    except Exception as e:
        print(f"Error reading {file_path}: {e}")
        return None

def create_data_digest_agent():
    tabular_data_tool = Tool(
        name="TabularDataTool",
        func=read_tabular_data,
        description="Reads and processes tabular data from various file formats including Excel and CSV."
    )
    
    return Agent(
        role='Data Digest Agent',
        goal='Digest and summarize company data from tabular data files',
        backstory="An AI agent responsible for securely reading and summarizing company tabular data, optimized for Excel files.",
        tools=[tabular_data_tool],
        verbose=True
    )

def create_analysis_agent(data_digest_crew):
    return Agent(
        role='Analysis Agent',
        goal='Analyze the summarized data to provide actionable insights',
        backstory="An AI agent that analyzes summarized tabular data to offer strategic improvement suggestions.",
        tools=[
            CrewAITool(
                name="Data Digest Crew",
                description="A crew that can digest and summarize tabular data, optimized for Excel",
                crew=data_digest_crew
            )
        ],
        verbose=True
    )

def create_recommendation_agent(analysis_crew):
    return Agent(
        role='Recommendation Agent',
        goal='Generate specific, actionable recommendations based on the analysis',
        backstory="An AI agent that takes analyzed data and creates targeted, implementable recommendations.",
        tools=[
            CrewAITool(
                name="Analysis Crew",
                description="A crew that can analyze summarized tabular data",
                crew=analysis_crew
            )
        ],
        verbose=True
    )

def create_data_digest_crew():
    agent = create_data_digest_agent()
    task = Task(
        description='Read and summarize the contents of the uploaded tabular data file, optimized for Excel.',
        expected_output='A summary of the company\'s state based on the tabular data.',
        agent=agent
    )
    return Crew(
        agents=[agent],
        tasks=[task],
        process=Process.sequential,
        verbose=True,
        memory=True
    )

def create_analysis_crew(data_digest_crew):
    agent = create_analysis_agent(data_digest_crew)
    task = Task(
        description='Analyze the summarized data to generate insights for company improvement.',
        expected_output='A comprehensive analysis report with key insights.',
        agent=agent
    )
    return Crew(
        agents=[agent],
        tasks=[task],
        process=Process.sequential,
        verbose=True,
        memory=True
    )

def setup_main_crew(data_digest_crew, analysis_crew):
    recommendation_agent = create_recommendation_agent(analysis_crew)
    task = Task(
        description='Generate specific, actionable recommendations based on the analysis.',
        expected_output='A list of targeted, implementable recommendations for company improvement.',
        agent=recommendation_agent
    )
    return Crew(
        agents=[recommendation_agent],
        tasks=[task],
        process=Process.sequential,
        verbose=True,
        memory=True
    )

def setup_crew():
    data_digest_crew = create_data_digest_crew()
    analysis_crew = create_analysis_crew(data_digest_crew)
    main_crew = setup_main_crew(data_digest_crew, analysis_crew)
    return main_crew

# Main execution
if __name__ == "__main__":
    crew = setup_crew()
    result = crew.kickoff()
    print(result)
