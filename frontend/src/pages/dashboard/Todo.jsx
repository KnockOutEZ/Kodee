/* eslint-disable react/button-has-type */
/* eslint-disable react/self-closing-comp */
import React, { useEffect, useState } from 'react';
import { DragDropContext, Droppable, Draggable } from 'react-beautiful-dnd';
import _ from 'lodash';
import { v4 } from 'uuid';
import '../../styles/todo.css';

function Todo() {
  const [inputState, setInputState] = useState([true, '']);
  const [text, setText] = useState('');
  const [state, setState] = useState({
    todo: {
      title: 'Todo',
      items: [],
    },
    'in-progress': {
      title: 'In Progress',
      items: [],
    },
    done: {
      title: 'Completed',
      items: [],
    },
  });

  useEffect(() => {
    const data = localStorage.getItem('todo');
    if (data) {
      setState(JSON.parse(data));
    }
  }, []);

  const handleDragEnd = ({ destination, source }) => {
    if (!destination) {
      return;
    }

    if (
      destination.index === source.index &&
      destination.droppableId === source.droppableId
    ) {
      return;
    }

    // Creating a copy of item before removing it from state
    const itemCopy = { ...state[source.droppableId].items[source.index] };

    setState((prev) => {
      prev = { ...prev };
      // Remove from previous items array
      prev[source.droppableId].items.splice(source.index, 1);

      // Adding to new items array location
      prev[destination.droppableId].items.splice(
        destination.index,
        0,
        itemCopy
      );

      return prev;
    });
    localStorage.setItem('todo', JSON.stringify(state));
  };

  const addItem = () => {
    if (!text) return;
    let  all
    setState((prev) => {
       all = {
        ...prev,
        todo: {
          title: 'Todo',
          items: [
            {
              id: v4(),
              name: text,
            },
            ...prev.todo.items,
          ],
        },
      };
      localStorage.setItem('todo', JSON.stringify(all));
      return all;
    });

    setText('');
    setInputState([true, '']);
  };

  const deleteItem = (id, droppableId) => {
    setState((prev) => {
      prev = { ...prev };
      prev[droppableId].items = prev[droppableId].items.filter(
        (item) => item.id !== id
      );
    localStorage.setItem('todo', JSON.stringify(prev));
      return prev;
    });
  };

  const clearCompleted = () => {
    setState((prev) => {
      prev = { ...prev };
      prev.done.items = [];
      localStorage.setItem('todo', JSON.stringify(prev));
      return prev;
    });
  };

  const editItem = (id, droppableId, value) => {
    setState((prev) => {
      prev = { ...prev };
      prev[droppableId].items = prev[droppableId].items.map((item) => {
        if (item.id === id) {
          item.name = value;
        }

        return item;
      });
      localStorage.setItem('todo', JSON.stringify(prev));
      return prev;
    });
  };

  return (
    <div className="todo">
      <div style={{ marginBottom: '20px' }}>
        <input
          className="add-input"
          type="text"
          value={text}
          onChange={(e) => {
            setText(e.target.value);
          }}
          onKeyDown={(e) =>
            e.key === 'Enter' &&
            addItem()
          }
        />
        <button type="button" className="top-button1 button" onClick={addItem}>
          Add
        </button>
        <button
          type="button"
          className="top-button2 button"
          onClick={clearCompleted}
        >
          Clear Completed
        </button>
      </div>
      <div className="todo-body">
        <DragDropContext onDragEnd={handleDragEnd}>
          {_.map(state, (data, key) => {
            return (
              <div key={key} className="column">
                <h3>{data.title}</h3>
                <Droppable droppableId={key}>
                  {(provided, snapshot) => {
                    return (
                      <div
                        ref={provided.innerRef}
                        {...provided.droppableProps}
                        className="droppable-col"
                      >
                        {data.items.map((el, index) => {
                          return (
                            <Draggable
                              key={el.id}
                              index={index}
                              draggableId={el.id}
                            >
                              {(provided, snapshot) => {
                                return (
                                  <div
                                    className={`item ${
                                      snapshot.isDragging && 'dragging'
                                    }`}
                                    ref={provided.innerRef}
                                    {...provided.draggableProps}
                                    {...provided.dragHandleProps}
                                  >
                                    {inputState[0] === false && inputState[1] === el.id ? (
                                      <div className="edit-delete">
                                      <input
                                        className="task-title-input"
                                        value={el.name}
                                        onChange={(e) =>
                                          editItem(el.id, key, e.target.value)
                                        }
                                        onKeyDown={(e) =>
                                          el.name !== "" && e.key === 'Enter' &&
                                          setInputState([true, ''])
                                        }
                                      />
                                      <button
                                        className="button deleteBtn"
                                        type="button"
                                        onClick={() => deleteItem(el.id, key)}
                                      >
                                        <i className="fa-solid fa-trash-can"></i>
                                      </button>
                                    </div>
                                      
                                    ) : (
                                      <button
                                        className="task-title"
                                        onClick={() => setInputState([false, el.id])}
                                      >
                                        {el.name}
                                      </button>
                                    )}
                                  </div>
                                );
                              }}
                            </Draggable>
                          );
                        })}
                        {provided.placeholder}
                      </div>
                    );
                  }}
                </Droppable>
              </div>
            );
          })}
        </DragDropContext>
      </div>
    </div>
  );
}

export default Todo;
