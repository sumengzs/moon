import { FC, Key, useEffect, useRef, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { Button, Form, Space } from 'antd'
import { DataOptionItem } from '@/components/Data/DataOption/DataOption.tsx'
import RouteBreadcrumb from '@/components/PromLayout/RouteBreadcrumb'
import { HeightLine, PaddingLine } from '@/components/HeightLine'
import { DataOption, DataTable, SearchForm } from '@/components/Data'
import { CopyOutlined } from '@ant-design/icons'
import { ActionKey } from '@/apis/data.ts'
import {
    StrategyItemType,
    defaultStrategyListRequest
} from '@/apis/home/monitor/strategy/types'
import { columns, searchItems, tableOperationItems } from './options'
import { Detail } from './child/Detail'
import strategyApi from '@/apis/home/monitor/strategy'
import { Status } from '@/apis/types'

const defaultPadding = 12

const Strategy: FC = () => {
    const navigate = useNavigate()
    const operationRef = useRef<HTMLDivElement>(null)
    const [queryForm] = Form.useForm()

    const [dataSource, setDataSource] = useState<StrategyItemType[]>([])

    const [loading, setLoading] = useState<boolean>(false)
    const [total, setTotal] = useState<number | string>(0)
    const [refresh, setRefresh] = useState<boolean>(false)
    const [openDetail, setOpenDetail] = useState<boolean>(false)
    const [operateId, setOperateId] = useState<number | undefined>()
    const [actionKey, setActionKey] = useState<ActionKey | undefined>(
        ActionKey.ADD
    )

    const handlerOpenDetail = (id?: number) => {
        setOperateId(id)
        setOpenDetail(true)
    }

    const handlerCloseDetail = () => {
        setOpenDetail(false)
    }

    // 获取数据
    const handlerGetData = () => {
        setLoading(true)
        strategyApi
            .getStrategyList(defaultStrategyListRequest)
            .then((res) => {
                setDataSource(res.list)
                setTotal(res.page.total)
            })
            .finally(() => {
                setLoading(false)
            })
    }

    useEffect(() => {
        handlerGetData()
    }, [refresh])

    // 刷新
    const handlerRefresh = () => {
        setRefresh((prev) => !prev)
    }

    // 分页变化
    const handlerTablePageChange = (page: number, pageSize?: number) => {
        console.log(page, pageSize)
    }

    // 可以批量操作的数据
    const handlerBatchData = (
        selectedRowKeys: Key[],
        selectedRows: StrategyItemType[]
    ) => {
        console.log(selectedRowKeys, selectedRows)
    }

    const toStrategyGroupPage = (record: StrategyItemType) => {
        console.log(record)
        navigate(`/home/monitor/strategy-group?id=${record.id}`)
    }

    const handlerBatchDelete = (ids: number[]) => {
        // TODO 批量删除
        console.log(ids)
    }

    // 批量修改状态
    const handlebatchChangeStatus = (ids: number[], status: Status) => {
        strategyApi.batchChangeStrategyStatus(ids, status).then(() => {
            handlerRefresh()
        })
    }

    // 处理表格操作栏的点击事件
    const handlerTableAction = (key: ActionKey, record: StrategyItemType) => {
        console.log(key, record)
        setActionKey(key)
        switch (key) {
            case ActionKey.STRATEGY_GROUP_LIST:
                toStrategyGroupPage(record)
                break
            case ActionKey.DETAIL:
                handlerOpenDetail(record.id)
                break
            case ActionKey.EDIT:
                handlerOpenDetail(record.id)
                break
            case ActionKey.DELETE:
                handlerBatchDelete([record.id])
                break
            case ActionKey.DISABLE:
                handlebatchChangeStatus([record.id], Status.STATUS_DISABLED)
                break
            case ActionKey.ENABLE:
                handlebatchChangeStatus([record.id], Status.STATUS_ENABLED)
                break
        }
    }

    const handlerDataOptionAction = (key: ActionKey) => {
        setActionKey(key)
        switch (key) {
            case ActionKey.ADD:
                handlerOpenDetail()
                break
        }
    }

    // 处理搜索表单的值变化
    const handlerSearFormValuesChange = (
        changedValues: any,
        allValues: any
    ) => {
        console.log(changedValues, allValues)
    }

    const leftOptions: DataOptionItem[] = [
        {
            key: ActionKey.BATCH_IMPORT,
            label: (
                <Button type="primary" loading={loading}>
                    批量导入
                </Button>
            )
        }
    ]

    const rightOptions: DataOptionItem[] = [
        {
            key: ActionKey.REFRESH,
            label: (
                <Button
                    type="primary"
                    loading={loading}
                    onClick={handlerRefresh}
                >
                    刷新
                </Button>
            )
        }
    ]

    return (
        <div className="bodyContent">
            <Detail
                open={openDetail}
                onClose={handlerCloseDetail}
                id={operateId}
                actionKey={actionKey}
                refresh={handlerRefresh}
            />
            <div ref={operationRef}>
                <RouteBreadcrumb />
                <HeightLine />
                <SearchForm
                    form={queryForm}
                    items={searchItems}
                    formProps={{
                        onValuesChange: handlerSearFormValuesChange
                    }}
                />
                <HeightLine />
                <DataOption
                    queryForm={queryForm}
                    rightOptions={rightOptions}
                    leftOptions={leftOptions}
                    action={handlerDataOptionAction}
                />
                <PaddingLine
                    padding={defaultPadding}
                    height={1}
                    borderRadius={4}
                />
            </div>

            <DataTable
                dataSource={dataSource}
                columns={columns}
                operationRef={operationRef}
                total={+total}
                loading={loading}
                operationItems={tableOperationItems}
                pageOnChange={handlerTablePageChange}
                rowSelection={{
                    onChange: handlerBatchData
                }}
                action={handlerTableAction}
                expandable={{
                    expandedRowRender: (record: StrategyItemType) => (
                        <Space>
                            <CopyOutlined />
                            <p style={{ margin: 0 }}>{record.expr}</p>
                        </Space>
                    )
                }}
            />
        </div>
    )
}

export default Strategy