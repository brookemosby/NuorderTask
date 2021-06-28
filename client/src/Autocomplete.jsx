import React, { Component, Fragment } from "react";

class Autocomplete extends Component {
  constructor(props) {
    super(props);

    this.state = {
      activeSuggestion: 0,
      suggestions: [],
      showSuggestions: false,
      userInput: ""
    };
  }

  onChange = e => {
    const userInput = e.currentTarget.value;

    let requestOptions = {
      method: 'GET'
    }

    let path = "http://localhost:8443/search/"+userInput;
    fetch(path, requestOptions)
    .then((response) => response.json())
    .then(result => {
      console.log(result);
      this.setState({
        activeSuggestion: 0,
        suggestions: result,
        showSuggestions: true,
        userInput: userInput
        });
    })
  };

  onClick = e => {
    this.setState({
      activeSuggestion: 0,
      suggestions: [],
      showSuggestions: false,
      userInput: e.currentTarget.innerText
    });
  };

  onKeyDown = e => {
    const { activeSuggestion, suggestions } = this.state;

    // User pressed the enter key
    if (e.keyCode === 13) {
      this.setState({
        activeSuggestion: 0,
        showSuggestions: false,
        userInput: suggestions[activeSuggestion].title
      });
    }
    // User pressed the up arrow
    else if (e.keyCode === 38) {
      if (activeSuggestion === 0) {
        return;
      }

      this.setState({ activeSuggestion: activeSuggestion - 1 });
    }
    // User pressed the down arrow
    else if (e.keyCode === 40) {
      if (activeSuggestion - 1 === suggestions.length) {
        return;
      }

      this.setState({ activeSuggestion: activeSuggestion + 1 });
    }
  };

  render() {
    const {
      onChange,
      onClick,
      onKeyDown,
      state: {
        activeSuggestion,
        suggestions,
        showSuggestions,
        userInput
      }
    } = this;

    let suggestionsListComponent;

    if (showSuggestions && userInput) {
      if (suggestions.length) {
        suggestionsListComponent = (
          <ul class="suggestions">
            {suggestions.map((suggestion, index) => {
              let className;

              // Flag the active suggestion with a class
              if (index === activeSuggestion) {
                className = "suggestion-active";
              }

              return (
                <li className={className} key={suggestion.title} onClick={onClick}>
                  {suggestion.title}
                  {suggestion.labels.map((label, index) => {
                    const {description, name } = label;
                    return (
                      <div>
                        <small
                          className='label-view'
                          key={index}
                          style = {{ ["borderRadius"]: 5 }}
                          title={description}
                        >
                          {name}
                        </small>
                      </div>
                    );
                  })}
                </li>
              );
            })}
          </ul>
        );
      } else {
        suggestionsListComponent = (
          <div class="no-suggestions">
            <em>No suggestions</em>
          </div>
        );
      }
    }

    return (
      <Fragment>
        <input
          type="text"
          onChange={onChange}
          onKeyDown={onKeyDown}
          value={userInput}
        />
        {suggestionsListComponent}
      </Fragment>
    );
  }
}

export default Autocomplete;
