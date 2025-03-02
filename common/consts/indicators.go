package consts

const Indicators = `
{
    "indicators": [
        {
            "display_name": "商品属性",
            "children": [
                {
                    "factor_code": "product_id",
                    "display_name": "商品ID",
                    "allow_operators": [
                        {
                            "operator_code": "in",
                            "display_name": "in",
                            "input_el_type": 1
                        }
                    ]
                },
                {
                    "factor_code": "product_name",
                    "display_name": "商品名",
                    "allow_operators": [
                        {
                            "operator_code": "contain_any",
                            "display_name": "contain any (auto analyze)",
                            "input_el_type": 2
                        }
                    ]
                },
                {
                    "factor_code": "product_status",
                    "display_name": "商品状态",
                    "allow_operators": [
                        {
                            "operator_code": "in",
                            "display_name": "in",
                            "input_el_type": 3,
                            "allow_values": [
                                {
                                    "display_name": "on sale",
                                    "value": "1"
                                },
                                {
                                    "display_name": "deleted by seller",
                                    "value": "2"
                                },
                                {
                                    "display_name": "deleted by platform",
                                    "value": "3"
                                }
                            ]
                        }
                    ]
                }
            ]
        }
    ]
}`
