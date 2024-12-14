import subprocess
from ollama import generate

# Function to extract app name
def extract_app_name(user_input):
    prompt = f"""
    You are an assistant that extracts app names from user requests.
    The user might say something like "Open Safari" or "Close Notes."
    Extract and return the name of the app they want to interact with. If it's unclear, reply with "Unknown."

    Examples:
    1. "Can you open Safari?" -> Safari
    2. "Close Notes, please." -> Notes
    3. "Launch Mail." -> Mail
    4. "Quit Calendar." -> Calendar
    5. "Just checking." -> Unknown

    User request: "{user_input}"
    Extracted app name:
    """

    result = generate(model="llama3.2", prompt=prompt).response.strip()
    print(f"[DEBUG] Raw app name response: {result}")  # Debugging raw response
    # Extract clean app name
    if "->" in result:
        app_name = result.split("->")[-1].strip().lower()
    else:
        app_name = result.split(":")[-1].strip().lower()
    print(f"[DEBUG] Extracted app name: {app_name}")
    return app_name

# Function to extract the operation (open or close)
def extract_operation(user_input):
    prompt = f"""
    You are an assistant that extracts operations from user requests.
    The user might say something like "Open Safari" or "Close Notes."
    Extract and return the operation they want to perform: either "open" or "close." 
    If it's unclear, reply with "Unknown."

    Examples:
    1. "Can you open Safari?" -> open
    2. "Close Notes, please." -> close
    3. "Launch Mail." -> open
    4. "Quit Calendar." -> close
    5. "Just checking." -> Unknown

    User request: "{user_input}"
    Extracted operation:
    """

    result = generate(model="llama3.2", prompt=prompt).response.strip()
    print(f"[DEBUG] Raw operation response: {result}")  # Debugging raw response
    
    # Extract the operation using a pattern match
    if "open" in result.lower():
        return "open"
    elif "close" in result.lower():
        return "close"
    return "unknown"  # Default to "unknown" if parsing fails

# Define a function to open macOS applications
def open_app(app_name):
    try:
        # Try to open from the Applications directory
        result = subprocess.run(["open", f"/Applications/{app_name.capitalize()}.app"], check=True)
        return f"Successfully opened {app_name.capitalize()} from /Applications."
    except subprocess.CalledProcessError:
        try:
            # Fallback to the System Applications directory
            result = subprocess.run(["open", f"/System/Applications/{app_name.capitalize()}.app"], check=True)
            return f"Successfully opened {app_name.capitalize()} from /System/Applications."
        except subprocess.CalledProcessError as e:
            return f"Failed to open {app_name}. Error: {str(e)}"

# Define a function to close macOS applications
def close_app(app_name):
    try:
        # Try to close the app using its name from both directories
        result = subprocess.run(["osascript", "-e", f'tell application "{app_name.capitalize()}" to quit'], check=True)
        return f"Successfully closed {app_name.capitalize()}."
    except subprocess.CalledProcessError as e:
        return f"Failed to close {app_name}. Error: {str(e)}"

# Process command function
def process_command(command):
    app_name = extract_app_name(command)
    operation = extract_operation(command)

    # Debugging extracted values
    print(f"[DEBUG] Final extracted values - app_name: {app_name}, operation: {operation}")

    if app_name == "unknown" or operation == "unknown":
        return "Unable to determine the app or operation from your request."

    if operation == "open":
        return open_app(app_name)
    elif operation == "close":
        return close_app(app_name)
    else:
        return f"Invalid operation. Extracted operation was '{operation}'."

# Main loop
if __name__ == "__main__":
    print("Welcome! I can help you open or close apps on your Mac. Type 'exit' to quit.")
    while True:
        user_input = input("What would you like me to do? ")
        if user_input.lower() == "exit":
            print("Goodbye!")
            break
        result = process_command(user_input)
        print(result)