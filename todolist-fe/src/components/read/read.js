import React, { useState, useEffect } from "react";
import { Table, Button, Checkbox } from 'semantic-ui-react';
import axios from "axios";
import { Link } from "react-router-dom";

export default function Read() {
    const [todolists, setTodolists] = useState([]);
    const [checked, setChecked] = useState(false);

    useEffect(() => {
        axios.get("http://localhost:8000/todolists")
            .then((getData) => {
                setTodolists(getData.data.data);
        })
    }, [])

    const getData = () => {
        axios.get("http://localhost:8000/todolists")
            .then((getData) => {
                setTodolists(getData.data.data);
            })
    }

    const onUpdate = (id, name, status) => {
        axios.put(`http://localhost:8000/todolist/${id}`, {
            name,
            status
        }).then(() => {
                getData();
            });
    }

    const onDelete = (id) => {
        axios.delete(`http://localhost:8000/todolist/${id}`)
            .then(() => {
                getData();
            })
    }

    return (
        <div>
            <Link to= '/create'>
                <Button primary>Create</Button>
            </Link>
            <Table celled textAlign='center'>
                <Table.Header>
                    <Table.Row>
                        <Table.HeaderCell>Todo Name</Table.HeaderCell>
                        <Table.HeaderCell>Todo Status</Table.HeaderCell>
                        <Table.HeaderCell colspan='2'>Todo Action</Table.HeaderCell>
                    </Table.Row>
                </Table.Header>
                <Table.Body>
                    { todolists.map((data) => {
                        return (
                            <Table.Row>
                                <Table.Cell>{data.Name}</Table.Cell>
                                <Table.Cell>{data.Status ? 'Completed' : 'Not Completed'}</Table.Cell>
                                <Table.Cell>
                                    <Checkbox checked={data.Status ? true : false} onClick={() => {onUpdate(data.ID, data.Name, !data.Status); setChecked(!checked); console.log(data.Status)}}></Checkbox>
                                </Table.Cell>
                                <Table.Cell>
                                    <Button negative onClick={() => onDelete(data.ID)}>Delete</Button>
                                </Table.Cell>
                            </Table.Row>
                        )    
                    })}
                </Table.Body>
            </Table>
        </div>
    )
}