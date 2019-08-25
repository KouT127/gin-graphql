import React from "react";
import {
    Button,
    Card,
    CardContent,
    CardHeader,
    CircularProgress,
    createStyles,
    makeStyles,
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableRow,
    Theme
} from "@material-ui/core";
import {GetTasksQueryVariables, TaskInput, useAddUserMutation, useGetTasksQuery, UserInput} from "../generated/graphql";
import {useSelector} from "react-redux";
import {userSelector} from "../reducers/UserReducer";

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        tasks: {
            display: `grid`,
            gridTemplateRows: `2000px`,
            gridTemplateColumns: `1fr 90% 1fr`,
        },
        mainSection: {
            gridColumn: `2 / 3`
        },
        card: {
            marginLeft: theme.spacing(1),
            marginRight: theme.spacing(1),
        },
        table: {
            minWidth: 650,
            marginBottom: 9
        },
    }),
);

interface Task {
    id: string,
    title: string,
    description: string,
    userName: string,
}


const TasksPage: React.FC = () => {
    const classes = useStyles();
    const userState = useSelector(userSelector);
    const isLoggedIn = !!userState.user;
    const variable: GetTasksQueryVariables = {first: 10, after: ""};
    const [addUser, {}] = useAddUserMutation();
    const {data, error, loading} = useGetTasksQuery({fetchPolicy: "cache-and-network", variables: variable},);

    if (loading || error) return <CircularProgress/>;

    const handleAdd = () => {
        const task: TaskInput = {title: 'web', description: 'desc', userId: '2'};
        const input: UserInput = {name: 'name', gender: 'men', tasks: [task, task]};
        const response = addUser({variables: {user: input}});
        response.then((response) => {
            console.log(response.data)
        })
    };
    const tasks = data!.tasks.edges.map((edge) => {
        const node = edge.node;
        const task: Task = {
            id: node.id,
            title: node.title,
            description: node.description,
            userName: node.user ? node.user.name : ''
        };
        return task
    });
    return (
        <div className={classes.tasks}>
            <div className={classes.mainSection}>
                {isLoggedIn ? (
                    <CardTable tasks={tasks}/>
                ) : (
                    <CircularProgress/>
                )}
                <Card>
                    <Button
                        variant="contained"
                        color="primary"
                        onClick={handleAdd}>
                        ADD
                    </Button>
                </Card>
            </div>
        </div>
    );
};

const CardTable = (props: { tasks: Array<Task>, }) => {
    const {tasks} = props;
    const classes = useStyles();
    return (
        <Card className={classes.card}>
            <CardHeader
                title={'Tasks'}
            />
            <CardContent>
                <Table className={classes.table}>
                    <TableHead>
                        <TableRow>
                            <TableCell>ID</TableCell>
                            <TableCell>title</TableCell>
                            <TableCell>description</TableCell>
                            <TableCell>user name</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {tasks.map((row, index) => (
                            <TableRow key={'tasks-table-row-' + index}>
                                <TableCell>{row.id}</TableCell>
                                <TableCell>{row.title}</TableCell>
                                <TableCell>{row.description}</TableCell>
                                <TableCell>{row.userName}</TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </CardContent>
        </Card>
    )
};

export default TasksPage;
